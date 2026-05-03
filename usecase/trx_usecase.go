package usecase

import (
    "errors"
    "fmt"
    "time"
    "mini-shop/domain"
)

type TrxUsecase interface {
    GetAllTrx(userID uint) ([]Trx, error)
    GetTrxByID(id, userID uint) (Trx, error)
    CreateTrx(trx Trx, details []DetailTrxInput, userID uint) (Trx, error)
}

type trxUsecase struct {
    trxRepo      domain.TrxRepository
    detailRepo   domain.DetailTrxRepository
    logRepo      domain.LogProdukRepository
    produkRepo   domain.ProdukRepository
    alamatRepo   domain.AlamatRepository
	
}

func NewTrxUsecase(
    trxRepo domain.TrxRepository,
    detailRepo domain.DetailTrxRepository,
    logRepo domain.LogProdukRepository,
    produkRepo domain.ProdukRepository,
    alamatRepo domain.AlamatRepository,
) domain.TrxUsecase {
    return &trxUsecase{
        trxRepo:    trxRepo,
        detailRepo: detailRepo,
        logRepo:    logRepo,
        produkRepo: produkRepo,
        alamatRepo: alamatRepo,
    }
}

func (u *trxUsecase) GetAllTrx(userID uint) ([]domain.Trx, error) {
    return u.trxRepo.FindAll(userID)
}

func (u *trxUsecase) GetTrxByID(id, userID uint) (domain.Trx, error) {
    trx, err := u.trxRepo.FindByID(id)
    if err != nil {
        return domain.Trx{}, errors.New("No Data Trx")
    }
    if trx.UserID != userID {
        return domain.Trx{}, errors.New("unauthorized")
    }
    return trx, nil
}

func (u *trxUsecase) CreateTrx(trx domain.Trx, details []domain.DetailTrxInput, userID uint) (domain.Trx, error) {
    // Verify alamat belongs to user
    alamat, err := u.alamatRepo.FindByID(trx.AlamatKirimID)
    if err != nil || alamat.UserID != userID {
        return domain.Trx{}, errors.New("alamat not found or unauthorized")
    }

    trx.UserID = userID
    trx.KodeInvoice = fmt.Sprintf("INV-%d", time.Now().Unix())
    trx.HargaTotal = 0

    // Create the trx first
    createdTrx, err := u.trxRepo.Create(trx)
    if err != nil {
        return domain.Trx{}, err
    }

    // Process each product
    for _, item := range details {
        produk, err := u.produkRepo.FindByID(item.ProductID)
        if err != nil {
            return domain.Trx{}, errors.New("product not found")
        }
        if produk.Stok < item.Kuantitas {
            return domain.Trx{}, errors.New("insufficient stock")
        }

        // Create log produk snapshot
        logProduk, err := u.logRepo.Create(domain.LogProduk{
            ProdukID:      produk.ID,
            NamaProduk:    produk.NamaProduk,
            Slug:          produk.Slug,
            HargaReseller: produk.HargaReseller,
            HargaKonsumen: produk.HargaKonsumen,
            Stok:          produk.Stok,
            Deskripsi:     produk.Deskripsi,
            TokoID:        produk.TokoID,
            CategoryID:    produk.CategoryID,
        })
        if err != nil {
            return domain.Trx{}, err
        }

        // Calculate total for this item
        itemTotal := produk.HargaKonsumen * item.Kuantitas
        createdTrx.HargaTotal += float64(itemTotal)

        // Create detail trx
        u.detailRepo.Create(domain.DetailTrx{
            TrxID:       createdTrx.ID,
            LogProdukID: logProduk.ID,
            TokoID:      produk.TokoID,
            Kuantitas:   item.Kuantitas,
            HargaTotal:  itemTotal,
        })

        // Deduct stock
        produk.Stok -= item.Kuantitas
        u.produkRepo.Update(produk)
    }

    // Update total price
    u.trxRepo.Update(createdTrx)
    return createdTrx, nil
}