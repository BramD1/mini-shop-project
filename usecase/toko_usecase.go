package usecase

import (
    "errors"
    "mini-shop/domain"
)

type tokoUsecase struct {
    tokoRepo domain.TokoRepository
}

func NewTokoUsecase(tokoRepo domain.TokoRepository) domain.TokoUsecase {
    return &tokoUsecase{tokoRepo: tokoRepo}
}

func (u *tokoUsecase) GetTokoByID(id uint) (domain.Toko, error) {
    toko, err := u.tokoRepo.FindByID(id)
    if err != nil {
        return domain.Toko{}, errors.New("Toko tidak ditemukan")
    }
    return toko, nil
}

func (u *tokoUsecase) GetTokoByUserID(userID uint) (domain.Toko, error) {
    return u.tokoRepo.FindByUserID(userID)
}

func (u *tokoUsecase) GetAllToko(limit, offset int, nama string) ([]domain.Toko, error) {
    return u.tokoRepo.FindAll(limit, offset, nama)
}

func (u *tokoUsecase) CreateToko(toko domain.Toko) (domain.Toko, error) {
    return u.tokoRepo.Create(toko)
}

func (u *tokoUsecase) UpdateToko(toko domain.Toko, userID uint) (domain.Toko, error) {
    existing, err := u.tokoRepo.FindByID(toko.ID)
    if err != nil {
        return domain.Toko{}, errors.New("Toko tidak ditemukan")
    }
    if existing.UserID != userID {
        return domain.Toko{}, errors.New("unauthorized")
    }
    return u.tokoRepo.Update(toko)
}

func (u *tokoUsecase) DeleteToko(id uint) error {
    return u.tokoRepo.Delete(id)
}