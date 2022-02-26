package services

import (
	"log"
	"net/http"
	"safebsc/features/alert_price/models"
	"safebsc/postgres"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func DeleteAlertPrice(c *gin.Context) {
	// log.Println(request.PathParameters)
	alertPriceId := c.Param("alertPriceId")
	if alertPriceId == "" {
		log.Println("AlertPriceId Invalid Empty")
		c.String(http.StatusBadRequest, "Error")
	}

	userId := c.Param("userId")
	if userId == "" {
		log.Println("UserId Invalid Empty")
		c.String(http.StatusBadRequest, "Error")
	}

	log.Printf("!!!!! DeleteAlertPrice id: %s by userId: %s\n", alertPriceId, userId)

	// connect db
	db := postgres.Connect()
	defer db.Close()

	var alertPrice []models.AlertPrice
	if err := db.Model(&alertPrice).
		Where("users_id = ?", userId).
		Where("is_active = ?", true).
		Select(); err != nil {
		log.Println("Where AlertPrice error", err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	if alertPrice == nil || len(alertPrice) == 0 {
		log.Println("AlertPriceItems not found")
		c.String(http.StatusNoContent, "Error")
	}

	for _, ct := range alertPrice {
		if strconv.Itoa(ct.ID) == alertPriceId && ct.IsActive {
			ct.IsActive = false
			ct.UpdatedDate = time.Now().UTC()
			ct.UpdatedBy = userId

			result, err := db.Model(&ct).WherePK().Update()
			if err != nil {
				log.Println("Update AlertPrice error", err.Error())
				c.String(http.StatusInternalServerError, err.Error())
			}

			if result.RowsAffected() > 0 {
				c.String(http.StatusOK, "Deleted")
			}
		}
	}

	c.String(http.StatusNotFound, "Not found")
}
