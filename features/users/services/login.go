package services

import (
	"encoding/json"
	"log"
	"net/http"
	"safebsc/features/postgres"
	"safebsc/features/users/constants"
	"safebsc/features/users/models"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	log.Printf("!!!!! Login")
	var loginRequest models.LoginRequest

	if err := c.BindJSON(&loginRequest); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error binding request model", "error": err})
		c.Abort()
		return
	}
	log.Printf("!!!!! Login request: %s\n", loginRequest)

	db := postgres.Connect()
	defer db.Close()

	var users []models.Users
	if err := db.Model(&users).Where("wallet_address = ?", loginRequest.WalletAddress).Select(); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error wallet address not found", "error": err})
		c.Abort()
		return
	}

	for _, u := range users {
		if u.WalletAddress == loginRequest.WalletAddress {
			var tps []models.TelegramProfile
			if err := db.Model(&tps).Where("users_id = ?", u.Id).Select(); err != nil {
				log.Printf("TelegramProfile failed : %s\n", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"message": "TelegramProfile failed not  found", "error": err})
				c.Abort()
				return
			}

			var telegramId int
			if tps != nil && len(tps) > 0 {
				log.Printf("TelegramProfile : %d", tps[0].TelegramID)
				telegramId = tps[0].TelegramID
			}

			response := models.LoginResponse{
				UsersId:    u.Id,
				TelegramID: telegramId,
			}

			contents, err := json.Marshal(response)
			if err != nil {
				log.Println(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to mapped content", "error": err})
				c.Abort()
				return
			}

			log.Printf("!!!!! LoginResponse contents: %s\n", string(contents))
			c.JSON(http.StatusOK, contents)
			c.Abort()
			return
		}
	}

	var commonDB = models.CommonDB{
		CreatedDate: time.Now(),
		CreatedBy:   "System",
		UpdatedDate: time.Now(),
		UpdatedBy:   "System",
	}

	// Users
	user := models.Users{
		Common:          loginRequest.Common,
		TelegramAccount: "",
		LineToken:       "",
		Level:           constants.LevelType.FREE,
		CommonDB:        commonDB,
	}

	if _, err := db.Model(&user).Insert(); err != nil {
		log.Printf("Insert User %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Insert User failed", "error": err})
		c.Abort()
		return
	}

	// Wallet Address
	walletAddress := models.WalletAddresses{
		Common:   loginRequest.Common,
		CommonDB: commonDB,
	}

	if _, err := db.Model(&walletAddress).Insert(); err != nil {
		log.Printf("Insert Wallet Address %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Insert Wallet Address Failed", "error": err})
		c.Abort()
		return
	}

	// Users Wallet Address
	userWalletAddress := models.UsersWalletAddresses{
		Common:            loginRequest.Common,
		UsersId:           user.Id,
		WalletAddressesId: walletAddress.Id,
		WalletType:        constants.WalletType.MAIN,
		CommonDB:          commonDB,
	}

	if _, err := db.Model(&userWalletAddress).Insert(); err != nil {
		log.Printf("Insert Users Wallet Address %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Insert Users Wallet Address Failed", "error": err})
		c.Abort()
		return
	}

	var tps []models.TelegramProfile
	if err := db.Model(&tps).Where("users_id = ?", user.Id).Select(); err != nil {
		log.Printf("TelegramProfile failed : %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "TelegramProfile failed", "error": err})
		c.Abort()
		return
	}

	var telegramId int
	if tps != nil && len(tps) > 0 {
		log.Printf("TelegramProfile : %d", tps[0].TelegramID)
		telegramId = tps[0].TelegramID
	}

	response := models.LoginResponse{
		UsersId:    user.Id,
		TelegramID: telegramId,
	}

	contents, err := json.Marshal(response)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Mapped response failed", "error": err})
		c.Abort()
		return
	}

	log.Printf("!!!!! LoginResponse contents: %s\n", string(contents))
	c.JSON(http.StatusCreated, contents)
	c.Abort()
	return
}
