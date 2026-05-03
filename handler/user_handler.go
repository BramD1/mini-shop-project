package handler

import (
    "net/http"
    "mini-shop/domain"
    "mini-shop/utils"
    "github.com/gin-gonic/gin"
)

type UserHandler struct {
    userUsecase domain.UserUsecase
}

func NewUserHandler(userUsecase domain.UserUsecase) *UserHandler {
    return &UserHandler{userUsecase: userUsecase}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
    // Get userID from JWT context, not URL
    userID, _ := c.Get("userID")

    user, err := h.userUsecase.GetUserByID(userID.(uint))
    if err != nil {
        c.JSON(http.StatusNotFound, utils.ErrorResponse("Failed to GET data", err.Error()))
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to GET data", user))
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
    // Get userID from JWT context
    userID, _ := c.Get("userID")

    var user domain.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to PUT data", err.Error()))
        return
    }

    // Set ID from JWT — user can only update themselves
    user.ID = userID.(uint)

    _, err := h.userUsecase.UpdateUser(user)
    if err != nil {
        c.JSON(http.StatusBadRequest, utils.ErrorResponse("Failed to PUT data", err.Error()))
        return
    }

    c.JSON(http.StatusOK, utils.SuccessResponse("Succeed to GET data", ""))
}