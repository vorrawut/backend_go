package server

import (
	"safebsc/controllers"
	alertPriceController "safebsc/features/alert_price/controllers"
	aprCurrentController "safebsc/features/apr_current/controllers"
	customTokenController "safebsc/features/custom_token/controllers"
	farmConfigController "safebsc/features/farm_config/controllers"
	favoriteTokenController "safebsc/features/favorite_token/controllers"
	initDataController "safebsc/features/init_data/controllers"
	myFarmController "safebsc/features/my_farm/controllers"
	notificationController "safebsc/features/notification/controllers"
	proAccountController "safebsc/features/pro_account/controllers"
	tokenPriceController "safebsc/features/token_price/controllers"
	tokenPriceGraphController "safebsc/features/token_price_graph/controllers"
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

		customTokenGroup := v1.Group("customToken")
		{
			customToken := new(customTokenController.CustomTokenController)
			customTokenGroup.GET("/custom-token", customToken.GetCustomToken)
			customTokenGroup.POST("/custom-token", customToken.AddCustomToken)
			customTokenGroup.DELETE("/custom-token", customToken.RemoveCustomToken)
		}

		notificationGroup := v1.Group("notification")
		{
			notification := new(notificationController.Controller)
			notificationGroup.GET("/notification", notification.SendNotification)
		}

		favoriteTokenGroup := v1.Group("favoriteToken")
		{
			favoriteToken := new(favoriteTokenController.Controller)
			favoriteTokenGroup.GET("/favorite-token", favoriteToken.GetFavoriteToken)
			favoriteTokenGroup.POST("/favorite-token", favoriteToken.AddFavoriteToken)
			favoriteTokenGroup.DELETE("/favorite-token", favoriteToken.RemoveFavoriteToken)
		}

		initDataGroup := v1.Group("initData")
		{
			initData := new(initDataController.Controller)
			initDataGroup.GET("/init-data", initData.InitDataHandler)
		}

		farmConfigGroup := v1.Group("farmConfig")
		{
			farmConfig := new(farmConfigController.Controller)
			farmConfigGroup.GET("/farm-config", farmConfig.GetFarmConfig)
			farmConfigGroup.POST("/farm-config", farmConfig.SetFarmConfig)
			farmConfigGroup.PUT("/farm-config/{address}", farmConfig.BulkUpdateFarmConfig)
		}

		aprCurrentGroup := v1.Group("aprCurrent")
		{
			aprCurrent := new(aprCurrentController.Controller)
			aprCurrentGroup.GET("/apr-current", aprCurrent.GetAprCurrent)
		}

		myFarmGroup := v1.Group("myFarm")
		{
			myFarm := new(myFarmController.Controller)
			myFarmGroup.GET("/my-farm/{walletAddress}", myFarm.GetMyFarm)
		}

		tokenPriceGroup := v1.Group("tokenPrice")
		{
			tokenPrice := new(tokenPriceController.Controller)
			tokenPriceGroup.GET("/token-price", tokenPrice.GetTokenPrice)
		}

		tokenPriceGraphGroup := v1.Group("tokenPriceGraph")
		{
			tokenPriceGraph := new(tokenPriceGraphController.Controller)
			tokenPriceGraphGroup.GET("/token-price-graph/{tokenAddress}", tokenPriceGraph.GetTokenPriceGraph)
		}

		proAccountGroup := v1.Group("proAccount")
		{
			proAccount := new(proAccountController.Controller)
			proAccountGroup.GET("/pro-account", proAccount.GetProAccount)
		}

	}

	return router

}
