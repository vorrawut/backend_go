package server

import (
	"safebsc/controllers"
	alertPriceController "safebsc/features/alert_price/controllers"
	"safebsc/features/sayhello"
	userController "safebsc/features/users/controllers"
	"safebsc/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	router.GET("/say", sayhello.Handler)

	router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(userController.UserController)
			userGroup.POST("/users", user.Login)
			userGroup.GET("/users/{userId}", user.GetUsersById)
		}

		alertPriceGroup := v1.Group("alertPrice")
		{
			alertPrice := new(alertPriceController.AlertPriceController)
			alertPriceGroup.GET("/alert-price/userId/{userId}/tokenAddress/", alertPrice.GetAlertPrice)
			alertPriceGroup.POST("/alert-price", alertPrice.AddAlertPrice)
			alertPriceGroup.DELETE("/alert-price/userId/{userId}/alertPriceId/{alertPriceId}", alertPrice.DeleteAlertPrice)
		}
	}

	return router

}
