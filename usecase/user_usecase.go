package usecase

import (
    "mini-shop/domain"
    "golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
    userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
    return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) GetUserByID(id uint) (domain.User, error) {
    return u.userRepo.FindByID(id)
}

func (u *userUsecase) GetUserByEmail(email string) (domain.User, error) {
    return u.userRepo.FindByEmail(email)
}

func (u *userUsecase) GetUserByNoTelp(noTelp string) (domain.User, error) {
    return u.userRepo.FindByNoTelp(noTelp)
}

func (u *userUsecase) CreateUser(user domain.User) (domain.User, error) {
    return u.userRepo.Create(user)
}

func (u *userUsecase) UpdateUser(user domain.User) (domain.User, error) {
    // Hash new password if provided
    if user.KataSandi != "" {
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.KataSandi), bcrypt.DefaultCost)
        if err != nil {
            return domain.User{}, err
        }
        user.KataSandi = string(hashedPassword)
    }
    return u.userRepo.Update(user)
}

func (u *userUsecase) DeleteUser(id uint) error {
    return u.userRepo.Delete(id)
}