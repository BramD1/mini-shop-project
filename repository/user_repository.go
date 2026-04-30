package repository

import (
    "mini-shop/domain"
    "gorm.io/gorm"
)

type userRepository struct {
    db *gorm.DB
}

// Constructor — creates a new instance of userRepository
func NewUserRepository(db *gorm.DB) domain.UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) FindByID(id uint) (domain.User, error) {
    var user domain.User
    err := r.db.First(&user, id).Error
    return user, err
}

func (r *userRepository) FindByEmail(email string) (domain.User, error) {
    var user domain.User
    err := r.db.Where("email = ?", email).First(&user).Error
    return user, err
}

func (r *userRepository) FindByNoTelp(noTelp string) (domain.User, error) {
    var user domain.User
    err := r.db.Where("no_telp = ?", noTelp).First(&user).Error
    return user, err
}

func (r *userRepository) Create(user domain.User) (domain.User, error) {
    err := r.db.Create(&user).Error
    return user, err
}

func (r *userRepository) Update(user domain.User) (domain.User, error) {
    err := r.db.Save(&user).Error
    return user, err
}

func (r *userRepository) Delete(id uint) error {
    err := r.db.Delete(&domain.User{}, id).Error
    return err
}