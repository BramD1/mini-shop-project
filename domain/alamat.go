package domain

import "gorm.io/gorm"

type Alamat struct {
    gorm.Model           // ← ID, CreatedAt, UpdatedAt, DeletedAt all included here
    UserID       uint   `json:"user_id"`
    JudulAlamat  string `json:"judul_alamat"`
    NamaPenerima string `json:"nama_penerima"`
    NoTelp       string `json:"no_telp"`
    DetailAlamat string `json:"detail_alamat"`
}

func (Alamat) TableName() string {
    return "alamat"
}

type AlamatRepository interface {
    FindByUserID(userID uint) ([]Alamat, error)
    FindByID(id uint) (Alamat, error)
    Create(alamat Alamat) (Alamat, error)
    Update(alamat Alamat) (Alamat, error)
    Delete(id uint) error
}

type AlamatUsecase interface {
    GetMyAlamat(userID uint) ([]Alamat, error)
    GetAlamatByID(id, userID uint) (Alamat, error)
    CreateAlamat(alamat Alamat) (Alamat, error)
    UpdateAlamat(alamat Alamat, userID uint) (Alamat, error)
    DeleteAlamat(id, userID uint) error
}