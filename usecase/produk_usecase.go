package usecase

import (
    "errors"
    "strings"
    "mini-shop/domain"
)

type produkUsecase struct {
    produkRepo domain.ProdukRepository
    tokoRepo   domain.TokoRepository
    fotoProdukRepo domain.FotoProdukRepository
}

func NewProdukUsecase(produkRepo domain.ProdukRepository, tokoRepo domain.TokoRepository, fotoProdukRepo domain.FotoProdukRepository) domain.ProdukUsecase {
    return &produkUsecase{
        produkRepo: produkRepo,
        tokoRepo:   tokoRepo,
        fotoProdukRepo: fotoProdukRepo,
    }
}

func (u *produkUsecase) GetAllProduk(limit, offset int, namaProduk string, categoryID, tokoID uint, maxHarga, minHarga int) ([]domain.Produk, error) {
    return u.produkRepo.FindAll(limit, offset, namaProduk, categoryID, tokoID, maxHarga, minHarga)
}

func (u *produkUsecase) GetProdukByID(id uint) (domain.Produk, error) {
    produk, err := u.produkRepo.FindByID(id)
    if err != nil {
        return domain.Produk{}, errors.New("No Data Product")
    }
    return produk, nil
}

func (u *produkUsecase) CreateProduk(produk domain.Produk) (domain.Produk, error) {
    // Generate slug
    produk.Slug = strings.ToLower(strings.ReplaceAll(produk.NamaProduk, " ", "-"))
    return u.produkRepo.Create(produk)
}

func (u *produkUsecase) UpdateProduk(produk domain.Produk, userID uint) (domain.Produk, error) {
    toko, err := u.tokoRepo.FindByUserID(userID)
    if err != nil {
        return domain.Produk{}, errors.New("unauthorized")
    }
    existing, err := u.produkRepo.FindByID(produk.ID)
    if err != nil {
        return domain.Produk{}, err
    }
    if existing.TokoID != toko.ID {
        return domain.Produk{}, errors.New("unauthorized")
    }
    if produk.NamaProduk != "" {
        produk.Slug = strings.ToLower(strings.ReplaceAll(produk.NamaProduk, " ", "-"))
    }
    return u.produkRepo.Update(produk)
}

func (u *produkUsecase) DeleteProduk(id, userID uint) error {
    toko, err := u.tokoRepo.FindByUserID(userID)
    if err != nil {
        return errors.New("unauthorized")
    }
    existing, err := u.produkRepo.FindByID(id)
    if err != nil {
        return errors.New("record not found")
    }
    if existing.TokoID != toko.ID {
        return errors.New("unauthorized")
    }
    return u.produkRepo.Delete(id)
}