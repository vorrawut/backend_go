package custom_token

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	models "safebsc/features/custom_token/models"

	"safebsc/postgres"

	"github.com/gin-gonic/gin"
)

type CustomTokenController struct{}

func (u CustomTokenController) GetCustomToken(c *gin.Context) {
	getCustomToken(c)
}

func (u CustomTokenController) AddCustomToken(c *gin.Context) {
	addCustomToken(c)
}

func (u CustomTokenController) RemoveCustomToken(c *gin.Context) {
	removeCustomToken(c)
}

func getCustomToken(c *gin.Context) {
	walletAddress := c.Query("walletAddress")
	if walletAddress == "" {
		log.Println("Missing wallet address parameter")
		c.String(http.StatusBadRequest, "Error")
	}

	db := postgres.Connect()
	defer db.Close()
	var customTokens = models.CustomTokens
	if err := db.Model(&customTokens).Where("wallet_address = ?", walletAddress).Select(); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	contents, err := json.MarshalIndent(customTokens, "", "\t")
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(http.StatusOK, string(contents))
}

func addCustomToken(c *gin.Context) {
	var customToken models.CustomToken
	jsonData, err1 := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(jsonData), &customToken)
	if err1 != nil || err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, "Error")
	}

	db := postgres.Connect()
	defer db.Close()
	var customTokens = models.CustomTokens
	if err := db.Model(&customTokens).Where("wallet_address = ?", customToken.WalletAddress).Select(); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	for _, ct := range customTokens {
		if ct.TokenAddress == customToken.TokenAddress {
			c.String(http.StatusBadRequest, err.Error())
		}
	}

	if _, err := db.Model(&customToken).Insert(); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(http.StatusCreated, "Created")
}

func removeCustomToken(c *gin.Context) {
	var customToken models.CustomToken
	jsonData, err1 := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(jsonData), &customToken)
	if err1 != nil || err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, "Error")
	}

	db := postgres.Connect()
	defer db.Close()
	result, err := db.Model(&customToken).Where(
		"wallet_address = ? and token_address = ?",
		customToken.WalletAddress,
		customToken.TokenAddress,
	).Delete()
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}
	if result.RowsAffected() > 0 {
		c.String(http.StatusOK, "Deleted")
	}
	c.String(http.StatusNotFound, "Not found")

}
