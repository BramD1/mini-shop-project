package middleware

import (
    "net/http"
    "mini-shop/utils"
    "github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        isAdmin, exists := c.Get("isAdmin")
        if !exists || !isAdmin.(bool) {
            c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("Failed to POST data", "Unauthorized"))
            return
        }
        c.Next()
    }
}