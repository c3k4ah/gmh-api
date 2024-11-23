package main

import (
	"os"

	router "github.com/c3k4ah/gmh-api/api/v1"
	_ "github.com/c3k4ah/gmh-api/docs"
	"github.com/c3k4ah/gmh-api/internal/configs"
)

// @title GMH API
// @version 1
// @description This is a sample server for GMH API.
func main() {

	configs.SetupDatabase()

	grouter := router.SetupRouter()

	grouter.Run(":" + os.Getenv("PORT"))

}
