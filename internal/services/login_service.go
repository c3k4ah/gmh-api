package services

import (
	"net/http"

	"github.com/c3k4ah/gmh-api/internal/configs"
	"github.com/c3k4ah/gmh-api/internal/models"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var user models.User

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := configs.Database.Where("username = ? AND password = ?", loginRequest.Username, loginRequest.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := GenerateToken(user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
