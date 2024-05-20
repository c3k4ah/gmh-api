package configs

import (
	"log"
	"os"

	// "github.com/go-sql-driver/mysql"
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
