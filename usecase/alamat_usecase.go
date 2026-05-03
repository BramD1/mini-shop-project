package usecase

import (
    "errors"
    "mini-shop/domain"
)

type alamatUsecase struct {
    alamatRepo domain.AlamatRepository
}

func NewAlamatUsecase(alamatRepo domain.AlamatRepository) domain.AlamatUsecase {
    return &alamatUsecase{alamatRepo: alamatRepo}
}

func (u *alamatUsecase) GetMyAlamat(userID uint) ([]domain.Alamat, error) {
    return u.alamatRepo.FindByUserID(userID)
}

func (u *alamatUsecase) GetAlamatByID(id, userID uint) (domain.Alamat, error) {
    alamat, err := u.alamatRepo.FindByID(id)
    if err != nil {
        return domain.Alamat{}, err
    }

    // Ownership check
    if alamat.UserID != userID {
        return domain.Alamat{}, errors.New("unauthorized")
    }

    return alamat, nil
}

func (u *alamatUsecase) CreateAlamat(alamat domain.Alamat) (domain.Alamat, error) {
    return u.alamatRepo.Create(alamat)
}

func (u *alamatUsecase) UpdateAlamat(alamat domain.Alamat, userID uint) (domain.Alamat, error) {
    // Check ownership first
    existing, err := u.alamatRepo.FindByID(alamat.ID)
    if err != nil {
        return domain.Alamat{}, err
    }

    if existing.UserID != userID {
        return domain.Alamat{}, errors.New("unauthorized")
    }

    return u.alamatRepo.Update(alamat)
}

func (u *alamatUsecase) DeleteAlamat(id, userID uint) error {
    // Check ownership first
    existing, err := u.alamatRepo.FindByID(id)
    if err != nil {
        return err
    }

    if existing.UserID != userID {
        return errors.New("unauthorized")
    }

    return u.alamatRepo.Delete(id)
}