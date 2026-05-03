package repository

import (
    "mini-shop/domain"
    "gorm.io/gorm"
)

type produkRepository struct {
    db *gorm.DB
}

func NewProdukRepository(db *gorm.DB) domain.ProdukRepository {
    return &produkRepository{db: db}
}

func (r *produkRepository) FindAll(limit, offset int, namaProduk string, categoryID, tokoID uint, maxHarga, minHarga int) ([]domain.Produk, error) {
    var produks []domain.Produk
    query := r.db.Preload("Toko").Preload("Category").Preload("FotosProduk")
    if namaProduk != "" {
        query = query.Where("nama_produk LIKE ?", "%"+namaProduk+"%")
    }
    if categoryID != 0 {
        query = query.Where("category_id = ?", categoryID)
    }
    if tokoID != 0 {
        query = query.Where("toko_id = ?", tokoID)
    }
    if minHarga != 0 {
        query = query.Where("harga_konsumen >= ?", minHarga)
    }
    if maxHarga != 0 {
        query = query.Where("harga_konsumen <= ?", maxHarga)
    }
    err := query.Limit(limit).Offset(offset).Find(&produks).Error
    return produks, err
}

func (r *produkRepository) FindByID(id uint) (domain.Produk, error) {
    var produk domain.Produk
    err := r.db.Preload("Toko").Preload("Category").Preload("FotosProduk").First(&produk, id).Error
    return produk, err
}

func (r *produkRepository) Create(produk domain.Produk) (domain.Produk, error) {
    err := r.db.Create(&produk).Error
    return produk, err
}

func (r *produkRepository) Update(produk domain.Produk) (domain.Produk, error) {
    err := r.db.Save(&produk).Error
    return produk, err
}

func (r *produkRepository) Delete(id uint) error {
    result := r.db.Delete(&domain.Produk{}, id)
    if result.RowsAffected == 0 {
        return gorm.ErrRecordNotFound
    }
    return result.Error
}