package profile

import (
	"encoding/json"
	"log"
	"net/http"
	"safebsc/postgres"

	"github.com/gin-gonic/gin"
)

func getProFile(c *gin.Context) {

	db := postgres.Connect()
	defer db.Close()

	walletAddress := c.Param("walletAddress")

	var user user
	if err := db.Model(&user).Where("token_address = ?", walletAddress).Select(); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	content, err := json.Marshal(user)
	if err != nil {
		Response(err.Error(), 500, c)
	}
	Response(string(content), 200, c)
}

func Response(body string, statusCode int, c *gin.Context) {
	c.String(statusCode, body)
}

func Handler(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		getProFile(c)
	default:
		c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}
