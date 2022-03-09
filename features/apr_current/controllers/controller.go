package controllers

import (
	"safebsc/features/apr_current/services"

	"github.com/gin-gonic/gin"
)

type AprController struct{}

func (u AprController) GetAprCurrent(c *gin.Context) {
	services.GetAprCurrent(c)
}
