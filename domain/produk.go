package domain

import "gorm.io/gorm"

type ProdukUsecase interface {
    GetAllProduk(limit, offset int, namaProduk string, categoryID, tokoID uint, maxHarga, minHarga int) ([]Produk, error)
    GetProdukByID(id uint) (Produk, error)
    CreateProduk(produk Produk, userID uint) (Produk, error)
    UpdateProduk(produk Produk, userID uint) (Produk, error)
    DeleteProduk(id, userID uint) error
}

type Produk struct {
	gorm.Model
	NamaProduk   string `json:"nama_produk"`
	Slug 	  string `json:"slug"`
	HargaReseller uint   `json:"harga_reseller"`
	HargaKonsumen uint   `json:"harga_konsumen"`
	Stok		uint   `json:"stok"`
	Deskripsi    string `json:"deskripsi"`
	TokoID     uint       `json:"toko_id"`
	CategoryID uint       `json:"category_id"`
	Toko       Toko       `gorm:"foreignKey:TokoID" json:"toko"`
	Category   Category   `gorm:"foreignKey:CategoryID" json:"category"`
	FotosProduk []FotoProduk `gorm:"foreignKey:ProdukID" json:"photos"`
}

func (Produk) TableName() string {
	return "produk"
}

type ProdukRepository interface {
	FindAll(limit, offset int, namaProduk string, categoryID, tokoID uint, maxHarga, minHarga int) ([]Produk, error)
	FindByID(id uint) (Produk, error)
	Create(produk Produk) (Produk, error)
	Update(produk Produk) (Produk, error)
	Delete(id uint) error
}

type ProdukUsecase interface {
	GetAllProduk(limit, offset int, namaProduk string, categoryID, tokoID uint, maxHarga, minHarga int) ([]Produk, error)
	GetProdukByID(id uint) (Produk, error)
	CreateProduk(produk Produk) (Produk, error)
	UpdateProduk(produk Produk) (Produk, error)
	DeleteProduk(id uint) error
}
