package main

import (
	"os"

	router "github.com/c3k4ah/gmh-api/api/v1"
	"github.com/c3k4ah/gmh-api/internal/configs"
)

func main() {

	configs.SetupDatabase()

	grouter := router.SetupRouter()

	grouter.Run(":" + os.Getenv("PORT"))

}
