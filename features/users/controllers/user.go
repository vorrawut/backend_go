package controllers

import (
	"safebsc/features/users/models"
	"safebsc/features/users/services"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userModel = new(models.Users)

func (u UserController) Login(c *gin.Context) {
	services.Login(c)
}

func (u UserController) GetUsersById(c *gin.Context) {
	services.GetUsersById(c)
}
