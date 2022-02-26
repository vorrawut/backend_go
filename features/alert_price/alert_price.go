package alert_price

import (
	"net/http"
	"safebsc/features/alert_price/services"

	"github.com/gin-gonic/gin"
)

func routers(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodPost:
		services.AddAlertPrice(c)
	case http.MethodGet:
		services.GetAlertPrice(c)
	case http.MethodDelete:
		services.DeleteAlertPrice(c)
	default:
		c.String(http.StatusMethodNotAllowed, "Error Method not allowed")
	}
}
