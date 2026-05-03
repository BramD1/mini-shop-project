package handler

import (
    "fmt"
    "net/http"
    "mini-shop/domain"
    "mini-shop/utils"
    "github.com/gin-gonic/gin"
)

type CategoryHandler struct {
    categoryUsecase domain.CategoryUsecase
}

func NewCategoryHandler(categoryUsecase domain.CategoryUsecase) *CategoryHandler {
    return &CategoryHandler{categoryUsecase: categoryUsecase}
}

func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
    categories, err := h.categoryUsecase.GetAllCategories()
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to GET data", err.Error()))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to GET data", categories))
}

func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
    var id uint
    fmt.Sscanf(c.Param("id"), "%d", &id)
    category, err := h.categoryUsecase.GetCategoryByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, utils.ErrorResponse("Failed to GET data", "No Data Category"))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to GET data", category))
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
    var category domain.Category
    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to POST data", err.Error()))
        return
    }
    created, err := h.categoryUsecase.CreateCategory(category)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to POST data", err.Error()))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to POST data", created.ID))
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
    var id uint
    fmt.Sscanf(c.Param("id"), "%d", &id)
    var category domain.Category
    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to PUT data", err.Error()))
        return
    }
    category.ID = id
    _, err := h.categoryUsecase.UpdateCategory(category)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to PUT data", err.Error()))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to GET data", ""))
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
    var id uint
    fmt.Sscanf(c.Param("id"), "%d", &id)
    err := h.categoryUsecase.DeleteCategory(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to GET data", err.Error()))
        return
    }
    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to GET data", ""))
}