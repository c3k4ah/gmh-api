package router

import (
	"github.com/c3k4ah/gmh-api/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/v1/appartement/:id", services.GetAppartement)

	return router
}
