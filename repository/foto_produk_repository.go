package repository

import (
    "mini-shop/domain"
    "gorm.io/gorm"
)

type fotoProdukRepository struct {
    db *gorm.DB
}

func NewFotoProdukRepository(db *gorm.DB) domain.FotoProdukRepository {
    return &fotoProdukRepository{db: db}
}

func (r *fotoProdukRepository) FindByProdukID(produkID uint) ([]domain.FotoProduk, error) {
    var fotos []domain.FotoProduk
    err := r.db.Where("produk_id = ?", produkID).Find(&fotos).Error
    return fotos, err
}

func (r *fotoProdukRepository) Create(foto domain.FotoProduk) (domain.FotoProduk, error) {
    err := r.db.Create(&foto).Error
    return foto, err
}

func (r *fotoProdukRepository) Update(foto domain.FotoProduk) (domain.FotoProduk, error) {
    err := r.db.Save(&foto).Error
    return foto, err
}

func (r *fotoProdukRepository) Delete(id uint) error {
    return r.db.Delete(&domain.FotoProduk{}, id).Error
}