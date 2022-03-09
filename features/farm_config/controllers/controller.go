package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"safebsc/postgres"

	models "safebsc/features/farm_config/models"

	"github.com/gin-gonic/gin"
)

type FarmConfigController struct{}

func (u FarmConfigController) GetFarmConfig(c *gin.Context) {
	getFarmConfig(c)
}

func (u FarmConfigController) SetFarmConfig(c *gin.Context) {
	setFarmConfig(c)
}

func (u FarmConfigController) BulkUpdateFarmConfig(c *gin.Context) {
	bulkUpdateFarmConfig(c)
}

type GetFarmConfigRequest struct {
	Address string `json:"email"`
}

func getFarmConfig(c *gin.Context) {
	walletAddress := c.Query("address")

	if !IsValidAddress(walletAddress) {
		Response("Invalid wallet address.", http.StatusBadRequest, c)
	}

	db := postgres.Connect()
	defer db.Close()
	// db.AutoMigrate(&models.FarmConfig{})

	fmt.Printf("db migration pass %s\n", walletAddress)

	var configs []models.FarmConfig
	if err := db.Model(&configs).Where("address = ?", walletAddress).Select(); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	fmt.Printf("query got %+v %s\n", configs, walletAddress)

	content, err := json.Marshal(configs)
	if err != nil {
		Response(err.Error(), http.StatusBadRequest, c)
	}

	c.String(http.StatusOK, string(content))

	Response(string(content), http.StatusOK, c)
}

func setFarmConfig(c *gin.Context) {

	var fc models.FarmConfig
	jsonData, err1 := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(jsonData), &fc)
	if err1 != nil || err != nil {
		Response(err.Error(), http.StatusInternalServerError, c)
	}

	if fc.FarmID == "" || fc.FarmName == "" {
		Response("FarmId, ChefAddress or FarmName Missing", http.StatusBadRequest, c)
	}

	db := postgres.Connect()
	defer db.Close()
	// db.AutoMigrate(&models.FarmConfig{})

	var configs []models.FarmConfig
	if err := db.Model(&configs).Where("address = ? and chef_address = ? and farm_id = ?", fc.Address, fc.ChefAddress, fc.FarmID).Select(); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	var content string

	if len(configs) > 0 {
		var configs = &models.FarmConfig{}
		if err := db.Model(&configs).Where("address = ? and chef_address = ? and farm_id = ?", fc.Address, fc.ChefAddress, fc.FarmID).Select(); err != nil {
			log.Println(err.Error())
			c.String(http.StatusInternalServerError, err.Error())
		}
		configs.IsDisable = fc.IsDisable
		db.Model(&configs).Update()
		content = "Updated"
	} else {
		nFc := &models.FarmConfig{
			Address:        fc.Address,
			FarmName:       fc.FarmName,
			TokenName:      fc.TokenName,
			FactoryAddress: fc.FactoryAddress,
			FarmID:         fc.FarmID,
			ChefAddress:    fc.ChefAddress,
			IsDisable:      fc.IsDisable,
		}
		db.Model(&nFc).Insert()
		content = "Created"
	}

	Response(content, http.StatusOK, c)
}

func bulkUpdateFarmConfig(c *gin.Context) {
	address := c.Param("address")
	if !IsValidAddress(address) {
		Response("invalid address", http.StatusBadRequest, c)
	}

	var fcs []models.FarmConfig
	jsonData, err1 := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(jsonData), &fcs)
	if err1 != nil || err != nil {
		Response(err.Error(), http.StatusInternalServerError, c)
	}

	filteredFcs := make([]models.FarmConfig, 0)

	for _, v := range fcs {
		if v.Address == address {
			filteredFcs = append(filteredFcs, v)
		}
	}

	if len(filteredFcs) == 0 {
		Response("nothing updated", http.StatusOK, c)
	}

	if len(filteredFcs) > 100 {
		Response("maximum update reach", http.StatusBadRequest, c)
	}

	db := postgres.Connect()
	defer db.Close()
	// db.AutoMigrate(&models.FarmConfig{})

	// TODO: check pro
	configsDB := make([]models.FarmConfig, 0)
	if err := db.Model(&configsDB).Where("address = ? ", address).Select(); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	toUpdateConfigsDB := make([]models.FarmConfig, 0)
	toCreateConfigsDB := make([]models.FarmConfig, 0)

	for _, v2 := range filteredFcs {
		isInDB := false
		for _, v := range configsDB {
			if v.Address == v2.Address && v.ChefAddress == v2.ChefAddress && v.FarmName == v2.FarmName && v.FarmID == v2.FarmID {
				toUpdateConfigsDB = append(toUpdateConfigsDB, v2)
				isInDB = true
				break
			}
		}
		if isInDB {
			continue
		}
		toCreateConfigsDB = append(toCreateConfigsDB, v2)
	}

	var content string
	if len(toUpdateConfigsDB) > 0 {
		for _, v := range toUpdateConfigsDB {
			var config = models.FarmConfig{}
			if err := db.Model(&config).Where("address = ? and chef_address = ? and farm_id = ?", v.Address, v.ChefAddress, v.FarmID).Select(); err != nil {
				log.Println(err.Error())
				c.String(http.StatusInternalServerError, err.Error())
			}
			config.IsDisable = v.IsDisable
			db.Model(&config).Update()
		}
		content += fmt.Sprintf("Updated : %d record", len(toUpdateConfigsDB))
	}

	if len(toCreateConfigsDB) > 0 {
		for _, v := range toCreateConfigsDB {
			nFc := &models.FarmConfig{
				Address:        v.Address,
				FarmName:       v.FarmName,
				TokenName:      v.TokenName,
				FactoryAddress: v.FactoryAddress,
				FarmID:         v.FarmID,
				ChefAddress:    v.ChefAddress,
				IsDisable:      v.IsDisable,
			}
			db.Model(&nFc).Insert()
		}
		content += fmt.Sprintf(", Created : %d record", len(toCreateConfigsDB))
	}

	Response(content, http.StatusOK, c)
}

func Response(body string, statusCode int, c *gin.Context) {
	c.String(statusCode, body)
}

func IsValidAddress(addr string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(addr)
}

func Handler(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		getFarmConfig(c)
	case "POST":
		setFarmConfig(c)
	case "PUT":
		bulkUpdateFarmConfig(c)

	default:
		c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}
