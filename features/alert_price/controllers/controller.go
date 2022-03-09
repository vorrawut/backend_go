package controllers

import (
	"safebsc/features/alert_price/services"

	"github.com/gin-gonic/gin"
)

type AlertPriceController struct{}

func (u AlertPriceController) AddAlertPrice(c *gin.Context) {
	services.AddAlertPrice(c)
}

func (u AlertPriceController) GetAlertPrice(c *gin.Context) {
	services.GetAlertPrice(c)
}

func (u AlertPriceController) DeleteAlertPrice(c *gin.Context) {
	services.DeleteAlertPrice(c)
}
