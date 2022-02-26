package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"safebsc/features/telegram_profile/models"
	"safebsc/postgres"
	"time"

	"github.com/gin-gonic/gin"
)

func AddTelegramProfile(c *gin.Context) {
	log.Printf("!!!!! AddTelegramProfile")
	var addTelegramProfileRequest models.AddTelegramProfileRequest
	jsonData, err1 := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(jsonData), &addTelegramProfileRequest)
	if err1 != nil || err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, "Error")
	}
	log.Printf("!!!!! AddTelegramProfileRequest request: %s\n", addTelegramProfileRequest)

	// Connect DB
	db := postgres.Connect()
	defer db.Close()

	profile := addTelegramProfileRequest.CommonTelegramProfile

	var tps []models.TelegramProfile
	if err := db.Model(&tps).Where("users_id = ?", profile.UsersID).Select(); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	for _, u := range tps {
		if u.UsersID == profile.UsersID {
			if u.IsActive == false {
				location, _ := time.LoadLocation("Asia/Bangkok")
				u.FirstName = profile.FirstName
				u.LastName = profile.LastName
				u.Username = profile.Username
				u.IsActive = true
				u.UpdatedDate = time.Now().In(location)

				if _, err := db.Model(&u).Where("users_id = ?", u.UsersID).Column("first_name", "last_name", "username", "is_active", "updated_date").Returning("*").Update(); err != nil {
					c.String(http.StatusInternalServerError, err.Error())
				}
			}

			c.String(http.StatusCreated, "Error")
		}
	}

	authDateParsed, err := time.Parse(time.RFC3339, addTelegramProfileRequest.AuthDate)
	if err != nil {
		fmt.Println(err)
	}

	ap := models.TelegramProfile{
		CommonTelegramProfile: profile,
		AuthDate:              authDateParsed,
		IsActive:              true,
		CreatedDate:           time.Now(),
		CreatedBy:             "System",
		UpdatedDate:           time.Now(),
		UpdatedBy:             "System",
	}

	if _, err := db.Model(&ap).Insert(); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(http.StatusCreated, "Created")
}
