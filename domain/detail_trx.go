package domain

import "gorm.io/gorm"

type DetailTrx struct {
	gorm.Model
	TrxID      uint   `json:"trx_id"`
	LogProdukID uint   `json:"id_log_produk"`
	TokoID      uint   `json:"id_toko"`
	Kuantitas   uint   `json:"kuantitas"`
	HargaTotal  uint   `json:"harga_total"`
}

func (DetailTrx) TableName() string {
	return "detail_trx"
}

type DetailTrxRepository interface {
	FindByTrxID(trxID uint) ([]DetailTrx, error)
	Create(detail DetailTrx) (DetailTrx, error)
	Update(detail DetailTrx) (DetailTrx, error)
	Delete(id uint) error
}

type DetailTrxUsecase interface {
	GetDetailsByTrxID(trxID uint) ([]DetailTrx, error)
	CreateDetail(detail DetailTrx) (DetailTrx, error)
	UpdateDetail(detail DetailTrx) (DetailTrx, error)
	DeleteDetail(id uint) error
}
