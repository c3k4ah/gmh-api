package main

import (
	"github.com/c3k4ah/gmh-api/internal/configs"
	"github.com/c3k4ah/gmh-api/internal/services"
)

func main() {
	configs.SetupDatabase()
	services.GetAppartement(3)
}
