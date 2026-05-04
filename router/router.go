package router

import (
    "mini-shop/domain"
    "mini-shop/handler"
    "mini-shop/middleware"
    "github.com/gin-gonic/gin"
)

func SetupRouter(
    authUsecase     domain.AuthUsecase,
    userUsecase     domain.UserUsecase,
    alamatUsecase   domain.AlamatUsecase,
    categoryUsecase domain.CategoryUsecase,
    produkUsecase   domain.ProdukUsecase,
    fotoProdukRepo  domain.FotoProdukRepository,
    trxUsecase      domain.TrxUsecase,
    tokoUsecase     domain.TokoUsecase,
) *gin.Engine {
    r := gin.Default()
    api := r.Group("/api/v1")

    // Auth
    authHandler := handler.NewAuthHandler(authUsecase)
    auth := api.Group("/auth")
    {
        auth.POST("/register", authHandler.Register)
        auth.POST("/login", authHandler.Login)
    }

    // User (protected)
    userHandler := handler.NewUserHandler(userUsecase)
    userGroup := api.Group("/user")
    userGroup.Use(middleware.AuthMiddleware())
    {
        userGroup.GET("", userHandler.GetProfile)
        userGroup.PUT("", userHandler.UpdateProfile)

        // Alamat
        alamatHandler := handler.NewAlamatHandler(alamatUsecase)
        alamat := userGroup.Group("/alamat")
        {
            alamat.GET("", alamatHandler.GetMyAlamat)
            alamat.GET("/:id", alamatHandler.GetAlamatByID)
            alamat.POST("", alamatHandler.CreateAlamat)
            alamat.PUT("/:id", alamatHandler.UpdateAlamat)
            alamat.DELETE("/:id", alamatHandler.DeleteAlamat)
        }
    }

    // Toko
    tokoHandler := handler.NewTokoHandler(tokoUsecase)
    toko := api.Group("/toko")
    {
        toko.GET("", tokoHandler.GetAllToko)
        toko.GET("/my", middleware.AuthMiddleware(), tokoHandler.GetMyToko)
        toko.GET("/:id", tokoHandler.GetTokoByID)
        toko.PUT("/:id", middleware.AuthMiddleware(), tokoHandler.UpdateToko)
    }

    // Category
    categoryHandler := handler.NewCategoryHandler(categoryUsecase)
    category := api.Group("/category")
    {
        category.GET("", categoryHandler.GetAllCategories)
        category.GET("/:id", categoryHandler.GetCategoryByID)
        category.POST("", middleware.AuthMiddleware(), middleware.AdminMiddleware(), categoryHandler.CreateCategory)
        category.PUT("/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), categoryHandler.UpdateCategory)
        category.DELETE("/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), categoryHandler.DeleteCategory)
    }

    // Product
    produkHandler := handler.NewProdukHandler(produkUsecase, fotoProdukRepo)
    product := api.Group("/product")
    {
        product.GET("", produkHandler.GetAllProduk)
        product.GET("/:id", produkHandler.GetProdukByID)
        product.POST("", middleware.AuthMiddleware(), produkHandler.CreateProduk)
        product.PUT("/:id", middleware.AuthMiddleware(), produkHandler.UpdateProduk)
        product.DELETE("/:id", middleware.AuthMiddleware(), produkHandler.DeleteProduk)
    }

    // Transaction
    trxHandler := handler.NewTrxHandler(trxUsecase)
    trx := api.Group("/trx")
    trx.Use(middleware.AuthMiddleware())
    {
        trx.GET("", trxHandler.GetAllTrx)
        trx.GET("/:id", trxHandler.GetTrxByID)
        trx.POST("", trxHandler.CreateTrx)
    }

    return r
}