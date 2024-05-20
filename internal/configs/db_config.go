package configs

import (
	"log"
	"os"

	"github.com/c3k4ah/gmh-api/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func DBConfig() mysql.Config {
	loadError := godotenv.Load("../../.env")
	if loadError != nil {
		log.Fatalf("Error .env file: %s", loadError)
	}

	cfg := mysql.Config{
		DSN: os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8&parseTime=True&loc=Local",
	}

	return cfg

}

func SetupDatabase() {
	var cfg mysql.Config = DBConfig()
	var err error

	Database, err = gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Database.AutoMigrate(models.Appartement{}, models.Localisation{})
}
