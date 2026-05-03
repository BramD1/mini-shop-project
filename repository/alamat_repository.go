package repository

import (
    "mini-shop/domain"
    "gorm.io/gorm"
)

type alamatRepository struct {
	db *gorm.DB
}

func NewAlamatRepository(db *gorm.DB) domain.AlamatRepository {
	return &alamatRepository{db: db}
}

func (r *alamatRepository) Create(alamat domain.Alamat) (domain.Alamat, error) {
	err := r.db.Create(&alamat).Error
	return alamat, err
}

func (r *alamatRepository) FindByID(id uint) (domain.Alamat, error) {
	var alamat domain.Alamat
	err := r.db.First(&alamat, id).Error
	return alamat, err
}

func (r *alamatRepository) Update(alamat domain.Alamat) (domain.Alamat, error) {
	var existing domain.Alamat
	err := r.db.First(&existing, alamat.ID).Error
	if err != nil {
		return domain.Alamat{}, err
	}
	err = r.db.Model(&existing).Updates(map[string]interface{}{
		"nama_penerima": alamat.NamaPenerima,
		"no_telp":       alamat.NoTelp,
		"alamat":        alamat.Alamat,
		"id_provinsi":   alamat.ProvinsiID,
		"id_kota":       alamat.KotaID,
	}).Error
	return existing, err
}