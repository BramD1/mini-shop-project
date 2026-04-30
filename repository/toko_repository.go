package repository

import (
    "mini-shop/domain"
    "gorm.io/gorm"
)

type tokoRepository struct {
    db *gorm.DB
}

func NewTokoRepository(db *gorm.DB) domain.TokoRepository {
    return &tokoRepository{db: db}
}

func (r *tokoRepository) FindAll(limit, offset int, nama string) ([]domain.Toko, error) {
    var tokos []domain.Toko
    query := r.db
    if nama != "" {
        query = query.Where("nama_toko LIKE ?", "%"+nama+"%")
    }
    err := query.Limit(limit).Offset(offset).Find(&tokos).Error
    return tokos, err
}

func (r *tokoRepository) FindByID(id uint) (domain.Toko, error) {
    var toko domain.Toko
    err := r.db.First(&toko, id).Error
    return toko, err
}

func (r *tokoRepository) FindByUserID(userID uint) (domain.Toko, error) {
    var toko domain.Toko
    err := r.db.Where("user_id = ?", userID).First(&toko).Error
    return toko, err
}

func (r *tokoRepository) Create(toko domain.Toko) (domain.Toko, error) {
    err := r.db.Create(&toko).Error
    return toko, err
}

func (r *tokoRepository) Update(toko domain.Toko) (domain.Toko, error) {
    err := r.db.Save(&toko).Error
    return toko, err
}

func (r *tokoRepository) Delete(id uint) error {
    err := r.db.Delete(&domain.Toko{}, id).Error
    return err
}