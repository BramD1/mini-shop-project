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
    var existing domain.FotoProduk
    err := r.db.First(&existing, foto.ID).Error
    if err != nil {
        return domain.FotoProduk{}, err
    }

    updateData := map[string]interface{}{}

    if foto.ProdukID != 0 {
        updateData["produk_id"] = foto.ProdukID
    }
    if foto.URL != "" {
        updateData["url"] = foto.URL
    }

    err = r.db.Model(&existing).Updates(updateData).Error
    return existing, err
}

func (r *fotoProdukRepository) Delete(id uint) error {
    return r.db.Delete(&domain.FotoProduk{}, id).Error
}