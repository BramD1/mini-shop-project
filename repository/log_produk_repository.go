package repository

import (
    "mini-shop/domain"
    "gorm.io/gorm"
)

type logProdukRepository struct {
    db *gorm.DB
}

func NewLogProdukRepository(db *gorm.DB) domain.LogProdukRepository {
    return &logProdukRepository{db: db}
}

func (r *logProdukRepository) FindByProdukID(produkID uint) ([]domain.LogProduk, error) {
    var logs []domain.LogProduk
    err := r.db.Where("produk_id = ?", produkID).Find(&logs).Error
    return logs, err
}

func (r *logProdukRepository) Create(log domain.LogProduk) (domain.LogProduk, error) {
    err := r.db.Create(&log).Error
    return log, err
}

func (r *logProdukRepository) Update(log domain.LogProduk) (domain.LogProduk, error) {
    err := r.db.Save(&log).Error
    return log, err
}

func (r *logProdukRepository) Delete(id uint) error {
    return r.db.Delete(&domain.LogProduk{}, id).Error
}