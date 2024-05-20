package services

import (
	"github.com/c3k4ah/gmh-api/internal/configs"
	"github.com/c3k4ah/gmh-api/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAppartement(ginCxt *gin.Context) {
	var (
		appartement models.Appartement
		result      gin.H
	)
	id := ginCxt.Param("id")
	error := configs.Database.Where("id = ?", id).Preload("Localisation").Find(&appartement).Error

	if error != nil {
		if gorm.ErrRecordNotFound == error {
			result = gin.H{
				"error": "Appartement not found",
			}
		} else {
			result = gin.H{
				"error": "Error while fetching appartement",
			}
		}
	} else {
		result = gin.H{
			"appartement": appartement,
		}
	}

	ginCxt.JSON(200, result)

}
