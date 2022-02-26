package server

import (
	"safebsc/controllers"
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
			userGroup.GET("/:id", user.GetUsersById)
		}
	}

	return router

}
