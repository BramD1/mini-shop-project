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

func (r *alamatRepository) FindByUserID(userID uint) ([]domain.Alamat, error) {
    var alamats []domain.Alamat
    err := r.db.Where("user_id = ?", userID).Find(&alamats).Error
    return alamats, err
}

func (r *alamatRepository) FindByID(id uint) (domain.Alamat, error) {
    var alamat domain.Alamat
    err := r.db.First(&alamat, id).Error
    return alamat, err
}

func (r *alamatRepository) Create(alamat domain.Alamat) (domain.Alamat, error) {
    err := r.db.Create(&alamat).Error
    return alamat, err
}

func (r *alamatRepository) Update(alamat domain.Alamat) (domain.Alamat, error) {
    var existing domain.Alamat
    err := r.db.First(&existing, alamat.ID).Error
    if err != nil {
        return domain.Alamat{}, err
    }

    err = r.db.Model(&existing).Updates(map[string]interface{}{
        "judul_alamat":  alamat.JudulAlamat,
        "nama_penerima": alamat.NamaPenerima,
        "no_telp":       alamat.NoTelp,
        "detail_alamat": alamat.DetailAlamat,
    }).Error
    return existing, err
}

func (r *alamatRepository) Delete(id uint) error {
    result := r.db.Delete(&domain.Alamat{}, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return gorm.ErrRecordNotFound
    }
    return nil
}