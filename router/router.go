package router

import (
    "mini-shop/domain"
    "mini-shop/handler"
    "github.com/gin-gonic/gin"
)

func SetupRouter(authUsecase domain.AuthUsecase) *gin.Engine {
    r := gin.Default()

    api := r.Group("/api/v1")

    // Auth routes
    authHandler := handler.NewAuthHandler(authUsecase)
    authGroup := api.Group("/auth")
    {
        authGroup.POST("/register", authHandler.Register)
        authGroup.POST("/login", authHandler.Login)
    }

    return r
}