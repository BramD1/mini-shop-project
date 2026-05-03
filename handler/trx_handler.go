package handler

import (
    "fmt"
    "net/http"
    "mini-shop/domain"
    "mini-shop/utils"
    "github.com/gin-gonic/gin"
)

type TrxHandler struct {
    trxUsecase domain.TrxUsecase
}

func NewTrxHandler(trxUsecase domain.TrxUsecase) *TrxHandler {
    return &TrxHandler{trxUsecase: trxUsecase}
}

func (h *TrxHandler) GetAllTrx(c *gin.Context) {
    userID, _ := c.Get("userID")
    trxs, err := h.trxUsecase.GetAllTrx(userID.(uint))
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to GET data", err.Error()))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to GET data", gin.H{
        "data":  trxs,
        "page":  0,
        "limit": 0,
    }))
}

func (h *TrxHandler) GetTrxByID(c *gin.Context) {
    userID, _ := c.Get("userID")
    var id uint
    fmt.Sscanf(c.Param("id"), "%d", &id)
    trx, err := h.trxUsecase.GetTrxByID(id, userID.(uint))
    if err != nil {
        c.JSON(http.StatusNotFound, utils.ErrorResponse("Failed to GET data", "No Data Trx"))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to GET data", trx))
}

func (h *TrxHandler) CreateTrx(c *gin.Context) {
    userID, _ := c.Get("userID")
    var input domain.TrxInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to POST data", err.Error()))
        return
    }

    trx := domain.Trx{
        MethodBayar:   input.MethodBayar,
        AlamatKirimID: input.AlamatKirim,
    }

    created, err := h.trxUsecase.CreateTrx(trx, input.DetailTrx, userID.(uint))
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to POST data", err.Error()))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to POST data", created.ID))
}