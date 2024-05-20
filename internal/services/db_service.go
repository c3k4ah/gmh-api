package services

import (
	"fmt"
	"log"

	"github.com/c3k4ah/gmh-api/internal/configs"
	"github.com/c3k4ah/gmh-api/internal/models"
	"gorm.io/gorm"
)

func GetAppartement(id uint) {

	appartement := models.Appartement{
		Model: gorm.Model{
			ID: id,
		},
	}

	result := configs.Database.Preload("Localisation").Find(&appartement)

	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println(appartement.Localisation.Adresse)

}

func CreateAppartement(appartement models.Appartement) {

	result := configs.Database.Create(&appartement)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

}
