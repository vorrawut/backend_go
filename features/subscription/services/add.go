package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"safebsc/features/subscription/models"
	"safebsc/postgres"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterSubscription(c *gin.Context) {

	db := postgres.Connect()
	defer db.Close()

	var requestData models.RegisterSubscriptionRequest

	jsonData, err1 := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(jsonData), &requestData)
	if err1 != nil || err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	var user models.Users
	var userSub []models.UsersSubscriptions
	var plan models.SubscriptionsPlans

	if err := db.Model(&user).Where("id = ?", requestData.UsersID).Select(); err != nil {
		log.Println("Query users error ", err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	if err := db.Model(&plan).Where("id = ?", requestData.SubscriptionsPlansID).Select(); err != nil {
		log.Println("Query plan error")
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	if err := db.Model(&userSub).Where("users_id = ?", user.ID).Select(); err != nil {
		log.Println("usersSubscription")
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	for _, u := range userSub {
		if u.IsActive {
			log.Printf("!!!!! already Created subscription")
			c.String(http.StatusOK, "Already subscription")
			return
		}
	}

	subscription := models.UsersSubscriptions{
		RegisterSubscriptionRequest: requestData,
		CreatedDate:                 time.Now().UTC(),
		CreatedBy:                   fmt.Sprint(user.ID),
		UpdatedDate:                 time.Now().UTC(),
		UpdatedBy:                   fmt.Sprint(user.ID),
	}

	if _, err := db.Model(&subscription).Insert(); err != nil {
		log.Println("subscription error")
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(http.StatusOK, "OK")
}
