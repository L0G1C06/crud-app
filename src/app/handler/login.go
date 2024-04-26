package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/L0G1C06/crud-app/schemas"
)

func LoginHandler(ctx *gin.Context) {
	var requestBody struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if requestBody.Username == "" || requestBody.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username and password must not be empty"})
		return
	}

	credentials := schemas.Credentials{}
	if err := db.Where("username = ? AND password = ?", requestBody.Username, requestBody.Password).First(&credentials).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "username or password are incorrect"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "login successful"})
}