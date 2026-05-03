package router

import (
    "mini-shop/domain"
    "mini-shop/handler"
    "mini-shop/middleware"
    "github.com/gin-gonic/gin"
)

func SetupRouter(authUsecase domain.AuthUsecase, userUsecase domain.UserUsecase) *gin.Engine {
    r := gin.Default()

    api := r.Group("/api/v1")

    // Auth routes
    authHandler := handler.NewAuthHandler(authUsecase)
    authGroup := api.Group("/auth")
    {
        authGroup.POST("/register", authHandler.Register)
        authGroup.POST("/login", authHandler.Login)
    }

    // User routes — protected
    userHandler := handler.NewUserHandler(userUsecase)
    userGroup := api.Group("/user")
    userGroup.Use(middleware.AuthMiddleware())
    {
        userGroup.GET("", userHandler.GetProfile)
        userGroup.PUT("", userHandler.UpdateProfile)
    }

    return r  // ← always last
}