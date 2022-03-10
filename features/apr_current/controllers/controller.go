package controllers

import (
	"safebsc/features/apr_current/services"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (u Controller) GetAprCurrent(c *gin.Context) {
	services.GetAprCurrent(c)
}
