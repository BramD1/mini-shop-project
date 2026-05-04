package domain

import "gorm.io/gorm"

type TrxUsecase interface {
    GetAllTrx(userID uint) ([]Trx, error)
    GetTrxByID(id, userID uint) (Trx, error)
    CreateTrx(trx Trx, details []DetailTrxInput, userID uint) (Trx, error)
}

type Trx struct {
	gorm.Model
	UserID   uint   `json:"id_user"`
	AlamatKirimID uint   `json:"alamat_kirim_id"`
	AlamatKirim   Alamat `gorm:"foreignKey:AlamatKirimID" json:"alamat_kirim"`
	HargaTotal float64 `json:"harga_total"`
	KodeInvoice string `json:"kode_invoice"`
	MethodBayar string `json:"method_bayar"`
	DetailTrx []DetailTrx `gorm:"foreignKey:TrxID" json:"detail_trx"`
}

func (Trx) TableName() string {
	return "trx"
}

type TrxRepository interface {
	FindByID(id uint) (Trx, error)
	FindByUserID(userID uint) ([]Trx, error)
	FindAll(userID uint) ([]Trx, error)
	Create(trx Trx) (Trx, error)
	Update(trx Trx) (Trx, error)
	Delete(id uint) error
}