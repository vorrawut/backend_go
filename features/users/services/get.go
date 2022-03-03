package services

import (
	"encoding/json"
	"log"
	"net/http"
	"safebsc/db"
	"safebsc/features/users/models"

	"github.com/gin-gonic/gin"
)

func GetUsersById(c *gin.Context) {
	userId := c.Param("userId")
	log.Printf("!!!!! GetUsersById: %s\n", userId)
	if userId == "" {
		log.Println("GetUsersById Invalid Empty")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "GetUsersById Invalid Empty"})
		c.Abort()
		return
	}

	// Connect DB
	db := db.GetConnection()

	var user []models.Users
	if err := db.Model(&user).Where("id = ?", userId).Select(); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "User not found", "error": err})
		c.Abort()
		return
	}

	if user == nil {
		log.Println("GetUsersById not found")
		c.JSON(http.StatusInternalServerError, gin.H{"message": "GetUsersById not found", "error": "err"})
		c.Abort()
		return
	}

	contents, err := json.Marshal(user)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Mapped model user", "error": err})
		c.Abort()
		return
	}

	log.Printf("!!!!! GetUsersById contents: %s\n", string(contents))
	c.JSON(http.StatusOK, contents)
	c.Abort()
	return
}
