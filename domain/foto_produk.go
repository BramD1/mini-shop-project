package domain

import "gorm.io/gorm"

type FotoProduk struct {
	gorm.Model
	ProdukID  uint   `json:"produk_id"`
	URL       string `json:"url"`
}

func (FotoProduk) TableName() string {
	return "foto_produk"
}

type FotoProdukRepository interface {
	FindByProdukID(produkID uint) ([]FotoProduk, error)
	Create(foto FotoProduk) (FotoProduk, error)
	Update(foto FotoProduk) (FotoProduk, error)
	Delete(id uint) error
}

type FotoProdukUsecase interface {
	GetFotosByProdukID(produkID uint) ([]FotoProduk, error)
	CreateFoto(foto FotoProduk) (FotoProduk, error)
	UpdateFoto(foto FotoProduk) (FotoProduk, error)
	DeleteFoto(id uint) error
}