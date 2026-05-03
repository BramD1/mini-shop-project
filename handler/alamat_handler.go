package handler

import (
    "fmt"
    "net/http"
    "mini-shop/domain"
    "mini-shop/utils"
    "github.com/gin-gonic/gin"
)

type AlamatHandler struct {
    alamatUsecase domain.AlamatUsecase
}

func NewAlamatHandler(alamatUsecase domain.AlamatUsecase) *AlamatHandler {
    return &AlamatHandler{alamatUsecase: alamatUsecase}
}

func (h *AlamatHandler) GetMyAlamat(c *gin.Context) {
    userID, _ := c.Get("userID")
    alamats, err := h.alamatUsecase.GetMyAlamat(userID.(uint))
    if err != nil {
        c.JSON(http.StatusNotFound, utils.ErrorResponse("Failed to GET data", err.Error()))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to GET data", alamats))
}

func (h *AlamatHandler) GetAlamatByID(c *gin.Context) {
    userID, _ := c.Get("userID")
    var id uint
    _, err := fmt.Sscanf(c.Param("id"), "%d", &id)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to GET data", "invalid ID"))
        return
    }

    alamat, err := h.alamatUsecase.GetAlamatByID(id, userID.(uint))
    if err != nil {
        c.JSON(http.StatusNotFound, utils.ErrorResponse("Failed to GET data", err.Error()))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to GET data", alamat))
}

func (h *AlamatHandler) CreateAlamat(c *gin.Context) {
    userID, _ := c.Get("userID")
    var alamat domain.Alamat
    if err := c.ShouldBindJSON(&alamat); err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to POST data", err.Error()))
        return
    }

    alamat.UserID = userID.(uint)
    createdAlamat, err := h.alamatUsecase.CreateAlamat(alamat)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to POST data", err.Error()))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to POST data", createdAlamat.ID))
}

func (h *AlamatHandler) UpdateAlamat(c *gin.Context) {
    userID, _ := c.Get("userID")
    var id uint
    _, err := fmt.Sscanf(c.Param("id"), "%d", &id)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to PUT data", "invalid ID"))
        return
    }

    var alamat domain.Alamat
    if err := c.ShouldBindJSON(&alamat); err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to PUT data", err.Error()))
        return
    }

    alamat.ID = id  // Set ID from URL param
    _, err = h.alamatUsecase.UpdateAlamat(alamat, userID.(uint))
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to GET data", err.Error()))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to GET data", ""))
}

func (h *AlamatHandler) DeleteAlamat(c *gin.Context) {
    userID, _ := c.Get("userID")
    var id uint
    _, err := fmt.Sscanf(c.Param("id"), "%d", &id)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to DELETE data", "invalid ID"))
        return
    }

    err = h.alamatUsecase.DeleteAlamat(id, userID.(uint))
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to GET data", err.Error()))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to GET data", ""))
}