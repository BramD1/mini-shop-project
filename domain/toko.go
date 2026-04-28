package domain

import "gorm.io/gorm"

type Toko struct {
	gorm.Model
	UserID   uint   `json:"id_user"`
	NamaToko string `json:"nama_toko"`
	URLFoto  string `json:"url_foto"`
}

func (Toko) TableName() string {
	return "toko"
}

type TokoRepository interface {
	FindAll(limit, offset int, nama string) ([]Toko, error)
	FindByUserID(userID uint) (Toko, error)
	FindByID(id uint) (Toko, error)
	Create(toko Toko) (Toko, error)
	Update(toko Toko) (Toko, error)
	Delete(id uint) error
}

type TokoUsecase interface {
	GetTokoByID(id uint) (Toko, error)
	GetTokoByUserID(userID uint) (Toko, error)
	CreateToko(toko Toko) (Toko, error)
	UpdateToko(toko Toko) (Toko, error)
	DeleteToko(id uint) error
}