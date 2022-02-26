package services

import (
	"encoding/json"
	"log"
	"net/http"
	"safebsc/features/subscription/models"
	"safebsc/postgres"

	"github.com/gin-gonic/gin"
)

func GetSubscriptionPlans(c *gin.Context) {
	log.Printf("!!!!! GetSubscriptionPlans")

	// Connect DB
	db := postgres.Connect()
	defer db.Close()

	var plans []models.SubscriptionsPlans
	if err := db.Model(&plans).Select(); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	if plans == nil || len(plans) == 0 {
		log.Println("Subscription plans not found")
		c.String(http.StatusNoContent, "Error")
	}

	contents, err := json.Marshal(plans)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	log.Printf("!!!!! GetSubscriptionPlans contents")
	c.String(http.StatusOK, string(contents))
}
