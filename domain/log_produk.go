package domain

import "gorm.io/gorm"


type LogProduk struct {
	gorm.Model
	ProdukID  uint   `json:"produk_id"`
	NamaProduk string `json:"nama_produk"`
	Slug string `json:"slug"`
	HargaReseller uint   `json:"harga_reseller"`
	HargaKonsumen uint   `json:"harga_konsumen"`
	Deskripsi  string `json:"deskripsi"`
	TokoID    uint   `json:"toko_id"`
	CategoryID uint   `json:"category_id"`
}

func (LogProduk) TableName() string {
	return "log_produk"
}

type LogProdukRepository interface {
	FindByProdukID(produkID uint) ([]LogProduk, error)
	Create(log LogProduk) (LogProduk, error)
	Update(log LogProduk) (LogProduk, error)
	Delete(id uint) error
}

type LogProdukUsecase interface {
	GetLogsByProdukID(produkID uint) ([]LogProduk, error)
	CreateLog(log LogProduk) (LogProduk, error)
	UpdateLog(log LogProduk) (LogProduk, error)
	DeleteLog(id uint) error
}