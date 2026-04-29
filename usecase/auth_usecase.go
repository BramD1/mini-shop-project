package usecase

import (
    "errors"
    "fmt"
    "mini-shop-project/domain"
    "mini-shop-project/utils"
    "golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
    userRepo domain.UserRepository
    tokoRepo domain.TokoRepository
}

func NewAuthUsecase(userRepo domain.UserRepository, tokoRepo domain.TokoRepository) domain.AuthUsecase {
    return &authUsecase{
        userRepo: userRepo,
        tokoRepo: tokoRepo,
    }
}

func (u *authUsecase) Register(user domain.User) error {
    // Step 1 — Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.KataSandi), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.KataSandi = string(hashedPassword)

    // Step 2 — Save the user
    createdUser, err := u.userRepo.Create(user)
    if err != nil {
        return err
    }

    // Step 3 — Auto create toko for the new user
    toko := domain.Toko{
        NamaToko: fmt.Sprintf("Toko-%s", createdUser.Nama),
        UserID:   createdUser.ID,
    }
    _, err = u.tokoRepo.Create(toko)
    if err != nil {
        return err
    }

    return nil
}

func (u *authUsecase) Login(noTelp, kataSandi string) (domain.User, error) {
    // Step 1 — Find user by noTelp
    user, err := u.userRepo.FindByNoTelp(noTelp)
    if err != nil {
        return domain.User{}, errors.New("No Telp atau kata sandi salah")
    }

    // Step 2 — Compare password
    err = bcrypt.CompareHashAndPassword([]byte(user.KataSandi), []byte(kataSandi))
    if err != nil {
        return domain.User{}, errors.New("No Telp atau kata sandi salah")
    }

    // Step 3 — Generate JWT token
    token, err := utils.GenerateToken(user.ID, user.Email, user.IsAdmin)
    if err != nil {
        return domain.User{}, err
    }

    // Step 4 — Attach token to user and return
    user.Token = token
    return user, nil
}