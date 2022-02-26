package subscription

import (
	"net/http"
	"safebsc/features/subscription/services"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodPost:
		services.RegisterSubscription(c)
		break
	case http.MethodGet:
		if c.Request.RequestURI == "/subscription/plans" {
			services.GetSubscriptionPlans(c)
		}
		break
	}
	c.String(http.StatusMethodNotAllowed, "Error")
}
