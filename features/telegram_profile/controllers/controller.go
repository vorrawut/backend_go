package controllers

import (
	"safebsc/features/telegram_profile/services"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (u Controller) AddTelegramProfile(c *gin.Context) {
	services.AddTelegramProfile(c)
}
