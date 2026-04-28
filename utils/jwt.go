package utils

import (
    "errors"
    "os"
    "time"
    "github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
    UserID  uint   `json:"id"`
    Email   string `json:"email"`
    IsAdmin bool   `json:"is_admin"`
    jwt.RegisteredClaims
}

func GenerateToken(userID uint, email string, isAdmin bool) (string, error) {
    secretKey := os.Getenv("JWT_SECRET")

    claims := JWTClaims{
        UserID:  userID,
        Email:   email,
        IsAdmin: isAdmin,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secretKey))
}

func ValidateToken(tokenString string) (*JWTClaims, error) {
    secretKey := os.Getenv("JWT_SECRET")

    token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return []byte(secretKey), nil
    })

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*JWTClaims)
    if !ok || !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}