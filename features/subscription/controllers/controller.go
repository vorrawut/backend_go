package controllers

import (
	"safebsc/features/subscription/services"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (u Controller) RegisterSubscription(c *gin.Context) {
	services.RegisterSubscription(c)
}

func (u Controller) GetSubscriptionPlans(c *gin.Context) {
	services.GetSubscriptionPlans(c)
}
