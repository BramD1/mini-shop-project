package repository

import (
    "mini-shop/domain"
    "gorm.io/gorm"
)

type userRepository struct {
    db *gorm.DB
}

// Constructor — creates a new instance of userRepository
func NewUserRepository(db *gorm.DB) domain.UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) FindByID(id uint) (domain.User, error) {
    var user domain.User
    err := r.db.First(&user, id).Error
    return user, err
}

func (r *userRepository) FindByEmail(email string) (domain.User, error) {
    var user domain.User
    err := r.db.Where("email = ?", email).First(&user).Error
    return user, err
}

func (r *userRepository) FindByNoTelp(noTelp string) (domain.User, error) {
    var user domain.User
    err := r.db.Where("no_telp = ?", noTelp).First(&user).Error
    return user, err
}

func (r *userRepository) Create(user domain.User) (domain.User, error) {
    err := r.db.Create(&user).Error
    return user, err
}

func (r *userRepository) Delete(id uint) error {
    err := r.db.Delete(&domain.User{}, id).Error
    return err
}

func (r *userRepository) Update(user domain.User) (domain.User, error) {
    // First find the existing user
    var existing domain.User
    err := r.db.First(&existing, user.ID).Error
    if err != nil {
        return domain.User{}, err
    }

    // Only update the fields that are provided
    err = r.db.Model(&existing).Updates(map[string]interface{}{
        "nama":          user.Nama,
        "no_telp":       user.NoTelp,
        "tanggal_lahir": user.TanggalLahir,
        "jenis_kelamin": user.JenisKelamin,
        "tentang":       user.Tentang,
        "pekerjaan":     user.Pekerjaan,
        "email":         user.Email,
        "kata_sandi":    user.KataSandi,
        "provinsi_id":   user.ProvinsiID,  // ← fixed
        "kota_id":       user.KotaID,      // ← fixed
    }).Error

    return existing, err
}