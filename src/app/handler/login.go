package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/L0G1C06/crud-app/schemas"
)

// @BasePath /api/v1

// @Summary Login User
// @Description Login user based on credentials saved on database
// @Tags User
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Request body"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/login [post]
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

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

	credentials := schemas.Signup{}
	if err := db.Where("username = ?", requestBody.Username).First(&credentials).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "username or password are incorrect"})
		return
	}

	if !CheckPasswordHash(requestBody.Password, credentials.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "username or password are incorrect"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "login successful"})
}
