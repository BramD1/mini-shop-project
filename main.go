package main

import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "mini-shop/config"
    "mini-shop/repository"
    "mini-shop/router"
    "mini-shop/usecase"
)

func main() {
    godotenv.Load()

    db, err := config.ConnectDB()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Repositories
    userRepo        := repository.NewUserRepository(db)
    tokoRepo        := repository.NewTokoRepository(db)
    alamatRepo      := repository.NewAlamatRepository(db)
    categoryRepo    := repository.NewCategoryRepository(db)
    produkRepo      := repository.NewProdukRepository(db)
    fotoProdukRepo  := repository.NewFotoProdukRepository(db)
    logProdukRepo   := repository.NewLogProdukRepository(db)
    trxRepo         := repository.NewTrxRepository(db)
    detailTrxRepo   := repository.NewDetailTrxRepository(db)

    // Usecases
    authUsecase     := usecase.NewAuthUsecase(userRepo, tokoRepo)
    userUsecase     := usecase.NewUserUsecase(userRepo)
    alamatUsecase   := usecase.NewAlamatUsecase(alamatRepo)
    categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
    produkUsecase   := usecase.NewProdukUsecase(produkRepo, tokoRepo, fotoProdukRepo)
    trxUsecase      := usecase.NewTrxUsecase(trxRepo, detailTrxRepo, logProdukRepo, produkRepo, alamatRepo)

    // Router
    r := router.SetupRouter(
        authUsecase,
        userUsecase,
        alamatUsecase,
        categoryUsecase,
        produkUsecase,
        fotoProdukRepo,
        trxUsecase,
    )
    r.Run(":" + os.Getenv("SERVER_PORT"))
}