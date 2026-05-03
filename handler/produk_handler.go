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

type ProdukHandler struct {
    produkUsecase  domain.ProdukUsecase
    fotoProdukRepo domain.FotoProdukRepository
}

func NewProdukHandler(produkUsecase domain.ProdukUsecase, fotoProdukRepo domain.FotoProdukRepository) *ProdukHandler {
    return &ProdukHandler{produkUsecase: produkUsecase, fotoProdukRepo: fotoProdukRepo}
}

func (h *ProdukHandler) GetAllProduk(c *gin.Context) {
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    offset := (page - 1) * limit
    namaProduk := c.Query("nama_produk")
    categoryID, _ := strconv.ParseUint(c.Query("category_id"), 10, 32)
    tokoID, _ := strconv.ParseUint(c.Query("toko_id"), 10, 32)
    maxHarga, _ := strconv.Atoi(c.Query("max_harga"))
    minHarga, _ := strconv.Atoi(c.Query("min_harga"))

    produks, err := h.produkUsecase.GetAllProduk(limit, offset, namaProduk, uint(categoryID), uint(tokoID), maxHarga, minHarga)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to GET data", err.Error()))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to GET data", gin.H{
        "data":  produks,
        "page":  page,
        "limit": limit,
    }))
}

func (h *ProdukHandler) GetProdukByID(c *gin.Context) {
    var id uint
    fmt.Sscanf(c.Param("id"), "%d", &id)
    produk, err := h.produkUsecase.GetProdukByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, utils.ErrorResponse("Failed to GET data", "No Data Product"))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to GET data", produk))
}

func (h *ProdukHandler) CreateProduk(c *gin.Context) {
    userID, _ := c.Get("userID")

    hargaReseller, _ := strconv.ParseUint(c.PostForm("harga_reseller"), 10, 32)
    hargaKonsumen, _ := strconv.ParseUint(c.PostForm("harga_konsumen"), 10, 32)
    stok, _ := strconv.ParseUint(c.PostForm("stok"), 10, 32)
    categoryID, _ := strconv.ParseUint(c.PostForm("category_id"), 10, 32)

    // Get toko by userID - handled in usecase
    produk := domain.Produk{
        NamaProduk:    c.PostForm("nama_produk"),
        HargaReseller: uint(hargaReseller),
        HargaKonsumen: uint(hargaKonsumen),
        Stok:          uint(stok),
        Deskripsi:     c.PostForm("deskripsi"),
        CategoryID:    uint(categoryID),
    }

    created, err := h.produkUsecase.CreateProduk(produk, userID.(uint))
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to POST data", err.Error()))
        return
    }

    // Handle multiple photo uploads
    form, _ := c.MultipartForm()
    if form != nil {
        files := form.File["photos"]
        os.MkdirAll("./uploads", os.ModePerm)
        for _, file := range files {
            filename := fmt.Sprintf("%d-%s", time.Now().Unix(), filepath.Base(file.Filename))
            if err := c.SaveUploadedFile(file, "./uploads/"+filename); err == nil {
                h.fotoProdukRepo.Create(domain.FotoProduk{
                    ProdukID: created.ID,
                    URL:      filename,
                })
            }
        }
    }

    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to POST data", created.ID))
}

func (h *ProdukHandler) UpdateProduk(c *gin.Context) {
    userID, _ := c.Get("userID")
    var id uint
    fmt.Sscanf(c.Param("id"), "%d", &id)

    hargaReseller, _ := strconv.ParseUint(c.PostForm("harga_reseller"), 10, 32)
    hargaKonsumen, _ := strconv.ParseUint(c.PostForm("harga_konsumen"), 10, 32)
    stok, _ := strconv.ParseUint(c.PostForm("stok"), 10, 32)
    categoryID, _ := strconv.ParseUint(c.PostForm("category_id"), 10, 32)

    produk := domain.Produk{
        NamaProduk:    c.PostForm("nama_produk"),
        HargaReseller: uint(hargaReseller),
        HargaKonsumen: uint(hargaKonsumen),
        Stok:          uint(stok),
        Deskripsi:     c.PostForm("deskripsi"),
        CategoryID:    uint(categoryID),
    }
    produk.ID = id

    _, err := h.produkUsecase.UpdateProduk(produk, userID.(uint))
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to PUT data", err.Error()))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to GET data", ""))
}

func (h *ProdukHandler) DeleteProduk(c *gin.Context) {
    userID, _ := c.Get("userID")
    var id uint
    fmt.Sscanf(c.Param("id"), "%d", &id)
    err := h.produkUsecase.DeleteProduk(id, userID.(uint))
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to GET data", err.Error()))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to GET data", ""))
}