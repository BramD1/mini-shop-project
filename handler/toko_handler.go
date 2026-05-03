package handler

import (
    "fmt"
    "net/http"
    "os"
    "path/filepath"
    "strconv"
    "time"
    "mini-shop/domain"
    "mini-shop/utils"
    "github.com/gin-gonic/gin"
)

type TokoHandler struct {
    tokoUsecase domain.TokoUsecase
}

func NewTokoHandler(tokoUsecase domain.TokoUsecase) *TokoHandler {
    return &TokoHandler{tokoUsecase: tokoUsecase}
}

func (h *TokoHandler) GetMyToko(c *gin.Context) {
    userID, _ := c.Get("userID")
    toko, err := h.tokoUsecase.GetTokoByUserID(userID.(uint))
    if err != nil {
        c.JSON(http.StatusNotFound, utils.ErrorResponse("Failed to GET data", err.Error()))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to GET data", toko))
}

func (h *TokoHandler) GetAllToko(c *gin.Context) {
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    nama := c.Query("nama")
    offset := (page - 1) * limit

    tokos, err := h.tokoUsecase.GetAllToko(limit, offset, nama)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to GET data", err.Error()))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to GET data", gin.H{
        "page":  page,
        "limit": limit,
        "data":  tokos,
    }))
}

func (h *TokoHandler) GetTokoByID(c *gin.Context) {
    var id uint
    fmt.Sscanf(c.Param("id"), "%d", &id)
    toko, err := h.tokoUsecase.GetTokoByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, utils.ErrorResponse("Failed to GET data", err.Error()))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to GET data", toko))
}

func (h *TokoHandler) UpdateToko(c *gin.Context) {
    userID, _ := c.Get("userID")
    var id uint
    fmt.Sscanf(c.Param("id"), "%d", &id)

    var toko domain.Toko
    toko.ID = id
    toko.NamaToko = c.PostForm("nama_toko")

    // Handle photo upload
    file, err := c.FormFile("photo")
    if err == nil {
        filename := fmt.Sprintf("%d-%s", time.Now().Unix(), filepath.Base(file.Filename))
        uploadPath := "./uploads/" + filename
        os.MkdirAll("./uploads", os.ModePerm)
        if err := c.SaveUploadedFile(file, uploadPath); err == nil {
            toko.URLFoto = filename
        }
    }

    _, err = h.tokoUsecase.UpdateToko(toko, userID.(uint))
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to UPDATE data", err.Error()))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to UPDATE data", "Update toko succeed"))
}