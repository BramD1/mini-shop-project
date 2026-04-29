package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama     string `json:"nama"`
	KataSandi string `json:"kata_sandi"`
	NoTelp string `gorm:"unique" json:"no_telp"`
	TanggalLahir string `json:"tanggal_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	Tentang string `json:"tentang"`
	Pekerjaan string `json:"pekerjaan"`
	Email string `gorm:"unique" json:"email"`
	ProvinsiID uint `json:"id_provinsi"`
	KotaID uint `json:"id_kota"`
	IsAdmin bool `json:"is_admin"`
	Token string `json:"token" gorm:"-"`
}

func (User) TableName() string {
	return "user"
}

type UserRepository interface {
	FindByID(id uint) (User, error)
	FindByEmail(email string) (User, error)
	FindByNoTelp(noTelp string) (User, error)
	Create(user User) (User, error)
	Update(user User) (User, error)
	Delete(id uint) error
}

type UserUsecase interface {
	GetUserByID(id uint) (User, error)
	GetUserByEmail(email string) (User, error)
	GetUserByNoTelp(noTelp string) (User, error)
	CreateUser(user User) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(id uint) error
}
