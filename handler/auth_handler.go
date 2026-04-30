package handler

import (
    "net/http"
    "mini-shop/domain"
    "mini-shop/utils"
    "github.com/gin-gonic/gin"
)

type AuthHandler struct {
    authUsecase domain.AuthUsecase
}

func NewAuthHandler(authUsecase domain.AuthUsecase) *AuthHandler {
    return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) Register(c *gin.Context) {
    var user domain.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to POST data", err.Error()))
        return
    }

    err := h.authUsecase.Register(user)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to POST data", err.Error()))
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to POST data", "Register Succeed"))
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req struct {
        NoTelp    string `json:"no_telp"`
        KataSandi string `json:"kata_sandi"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to POST data", err.Error()))
        return
    }

    user, err := h.authUsecase.Login(req.NoTelp, req.KataSandi)
    if err != nil {
        c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Failed to POST data", err.Error()))
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to POST data", user))
}