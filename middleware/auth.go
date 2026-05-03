package middleware

import (
    "net/http"
    "mini-shop/utils"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("token")
        if token == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("Unauthorized", "token is required"))
            return
        }

        claims, err := utils.ValidateToken(token)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("Unauthorized", err.Error()))
            return
        }

        // Inject user data into context for handlers to use
        c.Set("userID", claims.UserID)
        c.Set("isAdmin", claims.IsAdmin)
        c.Next()
    }
}