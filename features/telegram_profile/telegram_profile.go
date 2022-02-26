package telegram_profile

import (
	"net/http"
	"safebsc/features/telegram_profile/services"

	"github.com/gin-gonic/gin"
)

func routers(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodPost:
		services.AddTelegramProfile(c)
		break
	default:
		c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}
