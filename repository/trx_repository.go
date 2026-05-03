package repository

import (
    "mini-shop/domain"
    "gorm.io/gorm"
)

type trxRepository struct {
    db *gorm.DB
}

func NewTrxRepository(db *gorm.DB) domain.TrxRepository {
    return &trxRepository{db: db}
}

func (r *trxRepository) FindAll(userID uint) ([]domain.Trx, error) {
    var trxs []domain.Trx
    err := r.db.Preload("AlamatKirim").Preload("DetailTrx").Where("user_id = ?", userID).Find(&trxs).Error
    return trxs, err
}

func (r *trxRepository) FindByUserID(userID uint) ([]domain.Trx, error) {
    return r.FindAll(userID)
}

func (r *trxRepository) FindByID(id uint) (domain.Trx, error) {
    var trx domain.Trx
    err := r.db.Preload("AlamatKirim").Preload("DetailTrx").First(&trx, id).Error
    return trx, err
}

func (r *trxRepository) Create(trx domain.Trx) (domain.Trx, error) {
    err := r.db.Create(&trx).Error
    return trx, err
}

func (r *trxRepository) Update(trx domain.Trx) (domain.Trx, error) {
    err := r.db.Save(&trx).Error
    return trx, err
}

func (r *trxRepository) Delete(id uint) error {
    return r.db.Delete(&domain.Trx{}, id).Error
}