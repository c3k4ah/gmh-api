package router

import (
	"github.com/c3k4ah/gmh-api/internal/services"
	"github.com/c3k4ah/gmh-api/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()
	auth := router.Group("/")
	router.POST("/login", services.LoginHandler)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth.Use(middleware.AuthMiddleware())
	{
		router.GET("/v1/appartement/:id", services.GetAppartement)
		router.GET("/v1/appartements", services.GetAllAppartements)
		router.PUT("/v1/appartement/:id", services.UpdateAppartement)
		router.POST("/v1/appartement", services.CreateAppartement)
		router.DELETE("/v1/appartement/:id", services.CreateAppartement)
	}

	return router
}
