package main

import (
	"log"

	"github.com/c3k4ah/gmh-api/internal/configs"
	"github.com/c3k4ah/gmh-api/internal/models"
	"github.com/c3k4ah/gmh-api/internal/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var cfg mysql.Config = configs.DBConfig()
	var err error

	configs.Database, err = gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	configs.Database.AutoMigrate(models.Appartement{}, models.Localisation{})

	services.GetAppartement(3)
}
