package repository

import (
    "mini-shop/domain"
    "gorm.io/gorm"
)

type detailTrxRepository struct {
    db *gorm.DB
}

func NewDetailTrxRepository(db *gorm.DB) domain.DetailTrxRepository {
    return &detailTrxRepository{db: db}
}

func (r *detailTrxRepository) FindByTrxID(trxID uint) ([]domain.DetailTrx, error) {
    var details []domain.DetailTrx
    err := r.db.Where("trx_id = ?", trxID).Find(&details).Error
    return details, err
}

func (r *detailTrxRepository) Create(detail domain.DetailTrx) (domain.DetailTrx, error) {
    err := r.db.Create(&detail).Error
    return detail, err
}

func (r *detailTrxRepository) Update(detail domain.DetailTrx) (domain.DetailTrx, error) {
    err := r.db.Save(&detail).Error
    return detail, err
}

func (r *detailTrxRepository) Delete(id uint) error {
    return r.db.Delete(&domain.DetailTrx{}, id).Error
}