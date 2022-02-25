package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"safebsc/db"
	"safebsc/features/users/models"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userModel = new(models.Users)

// func (u UserController) Login(c *gin.Context) {
// 	if c.Param("id") != "" {
// 		user, err := userModel.GetByID(c.Param("id"))
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
// 			c.Abort()
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"message": "User founded!", "user": user})
// 		return
// 	}
// 	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
// 	c.Abort()
// 	return
// }

func (u UserController) GetUsersById(c *gin.Context) {
	userId := c.Param("userId")
	log.Printf("!!!!! GetUsersById: %s\n", userId)
	if userId == "" {
		log.Println("GetUsersById Invalid Empty")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve userId"})
		c.Abort()
		return
	}

	// Connect DB
	db := db.GetConnection()

	var user []models.Users
	if err := db.Model(&user).Where("id = ?", userId).Select(); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}

	if user == nil {
		log.Println("GetUsersById not found")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": "err"})
		c.Abort()
		return
	}

	contents, err := json.Marshal(user)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}

	log.Printf("!!!!! GetUsersById contents: %s\n", string(contents))
	c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
	c.Abort()
	return
}
