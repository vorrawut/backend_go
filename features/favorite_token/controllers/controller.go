package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"safebsc/features/favorite_token/models"
	"safebsc/postgres"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (u Controller) GetFavoriteToken(c *gin.Context) {
	getFavoriteToken(c)
}

func (u Controller) AddFavoriteToken(c *gin.Context) {
	addFavoriteToken(c)
}

func (u Controller) RemoveFavoriteToken(c *gin.Context) {
	removeFavoriteToken(c)
}

func getFavoriteToken(c *gin.Context) {
	walletAddress := c.Query("walletAddress")
	if walletAddress == "" {
		log.Println("Missing wallet address parameter")
		c.String(http.StatusBadRequest, "Error")
	}

	db := postgres.Connect()
	defer db.Close()
	var favoriteTokens = models.FavoriteTokens
	if err := db.Model(&favoriteTokens).Where("wallet_address = ?", walletAddress).Select(); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	contents, err := json.MarshalIndent(favoriteTokens, "", "\t")
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(http.StatusOK, string(contents))
}

func addFavoriteToken(c *gin.Context) {
	var favoriteToken models.FavoriteToken
	jsonData, err1 := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(jsonData), &favoriteToken)
	if err1 != nil || err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, "Error")
	}

	db := postgres.Connect()
	defer db.Close()
	var favoriteTokens = models.FavoriteTokens
	if err := db.Model(&favoriteTokens).Where("wallet_address = ?", favoriteToken.WalletAddress).Select(); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	for _, ft := range favoriteTokens {
		if ft.TokenAddress == favoriteToken.TokenAddress {
			c.String(http.StatusBadRequest, err.Error())
		}
	}

	if _, err := db.Model(&favoriteToken).Insert(); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(http.StatusCreated, "Created")
}

func removeFavoriteToken(c *gin.Context) {
	var favoriteToken models.FavoriteToken
	jsonData, err1 := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(jsonData), &favoriteToken)
	if err1 != nil || err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, "Error")
	}

	db := postgres.Connect()
	defer db.Close()
	db.Model(&favoriteToken).Where(
		"wallet_address = ? and token_address = ?",
		favoriteToken.WalletAddress,
		favoriteToken.TokenAddress,
	).Delete()
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(http.StatusOK, "Deleted")
}

func Handler(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		getFavoriteToken(c)
	case http.MethodPost:
		addFavoriteToken(c)
	case http.MethodDelete:
		removeFavoriteToken(c)
	default:
		c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}
