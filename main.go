package main

import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "mini-shop/config"
    "mini-shop/repository"
    "mini-shop/usecase"
    "mini-shop/router"
)

func main() {
    godotenv.Load()

    db, err := config.ConnectDB()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Repositories
    userRepo := repository.NewUserRepository(db)
    tokoRepo := repository.NewTokoRepository(db)

    // Usecases
    authUsecase := usecase.NewAuthUsecase(userRepo, tokoRepo)
    userUsecase := usecase.NewUserUsecase(userRepo)  // ← moved before Run

    // Router
    r := router.SetupRouter(authUsecase, userUsecase)  // ← pass both
    r.Run(":" + os.Getenv("SERVER_PORT"))  // ← always last
}