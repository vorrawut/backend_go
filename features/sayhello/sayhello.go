package sayhello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}
