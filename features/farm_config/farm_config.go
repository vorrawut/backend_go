package farm_config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Base struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type FarmConfig struct {
	Base
	Address        string `json:"address" bson:"address" binding:"required" `
	FarmName       string `json:"farmName" bson:"farm_name" binding:"required"`
	TokenName      string `json:"tokenName" bson:"token_name"`
	FactoryAddress string `json:"factoryAddress" bson:"factory_address"`
	FarmID         string `json:"farmId" bson:"farm_id" binding:"required"`
	ChefAddress    string `json:"chefAddress" bson:"chef_address"`
	IsDisable      *bool  `json:"isDisable" bson:"is_disable" binding:"required" pg:",use_zero"`
}

type GetFarmConfigRequest struct {
	Address string `json:"email"`
}

func getFarmConfig(c *gin.Context) {
	walletAddress := c.Query("address")

	if !IsValidAddress(walletAddress) {
		Response("Invalid wallet address.", http.StatusBadRequest, c)
	}

	postgresConnector := PostgresConnector{}
	db, err := postgresConnector.GetConnection()
	if err != nil {
		Response(err.Error(), http.StatusBadRequest, c)
	}
	defer db.Close()
	db.AutoMigrate(&FarmConfig{})

	fmt.Printf("db migration pass %s\n", walletAddress)

	farmConfig := &FarmConfig{Address: walletAddress}

	var configs []FarmConfig
	db.Where(farmConfig).Find(&configs)

	fmt.Printf("query got %+v %s\n", configs, walletAddress)

	content, err := json.Marshal(configs)
	if err != nil {
		Response(err.Error(), http.StatusBadRequest, c)
	}

	c.String(http.StatusOK, string(content))

	Response(string(content), http.StatusOK, c)
}

func setFarmConfig(c *gin.Context) {

	var fc FarmConfig
	jsonData, err1 := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(jsonData), &fc)
	if err1 != nil || err != nil {
		Response(err.Error(), http.StatusInternalServerError, c)
	}

	if fc.FarmID == "" || fc.FarmName == "" {
		Response("FarmId, ChefAddress or FarmName Missing", http.StatusBadRequest, c)
	}

	postgresConnector := PostgresConnector{}
	db, err := postgresConnector.GetConnection()
	if err != nil {
		Response(err.Error(), http.StatusInternalServerError, c)
	}
	defer db.Close()
	db.AutoMigrate(&FarmConfig{})

	var configs []FarmConfig
	db.Where(&FarmConfig{Address: fc.Address, ChefAddress: fc.ChefAddress, FarmID: fc.FarmID}).Find(&configs)
	var content string

	if len(configs) > 0 {
		db.Model(&FarmConfig{}).Where(&FarmConfig{Address: fc.Address, ChefAddress: fc.ChefAddress, FarmID: fc.FarmID}).Select("is_disable").Updates(map[string]interface{}{"is_disable": *fc.IsDisable})
		content = "Updated"
	} else {
		nFc := &FarmConfig{
			Address:        fc.Address,
			FarmName:       fc.FarmName,
			TokenName:      fc.TokenName,
			FactoryAddress: fc.FactoryAddress,
			FarmID:         fc.FarmID,
			ChefAddress:    fc.ChefAddress,
			IsDisable:      fc.IsDisable,
		}
		db.Create(&nFc)
		content = "Created"
	}

	Response(content, http.StatusOK, c)
}

func bulkUpdateFarmConfig(c *gin.Context) {
	address := c.Param("address")
	if !IsValidAddress(address) {
		Response("invalid address", http.StatusBadRequest, c)
	}

	var fcs []FarmConfig
	jsonData, err1 := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(jsonData), &fcs)
	if err1 != nil || err != nil {
		Response(err.Error(), http.StatusInternalServerError, c)
	}

	filteredFcs := make([]FarmConfig, 0)

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

	postgresConnector := PostgresConnector{}
	db, err := postgresConnector.GetConnection()
	if err != nil {
		Response(err.Error(), http.StatusInternalServerError, c)
	}
	defer db.Close()
	db.AutoMigrate(&FarmConfig{})

	// TODO: check pro

	configsDB := make([]FarmConfig, 0)
	db.Where(&FarmConfig{Address: address}).Find(&configsDB)

	toUpdateConfigsDB := make([]FarmConfig, 0)
	toCreateConfigsDB := make([]FarmConfig, 0)

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
			db.Model(&FarmConfig{}).Where(&FarmConfig{Address: v.Address, ChefAddress: v.ChefAddress, FarmID: v.FarmID}).Select("is_disable").Updates(map[string]interface{}{"is_disable": *v.IsDisable})
		}
		content += fmt.Sprintf("Updated : %d record", len(toUpdateConfigsDB))
	}

	if len(toCreateConfigsDB) > 0 {
		for _, v := range toCreateConfigsDB {
			nFc := &FarmConfig{
				Address:        v.Address,
				FarmName:       v.FarmName,
				TokenName:      v.TokenName,
				FactoryAddress: v.FactoryAddress,
				FarmID:         v.FarmID,
				ChefAddress:    v.ChefAddress,
				IsDisable:      v.IsDisable,
			}
			db.Create(&nFc)
		}
		content += fmt.Sprintf(", Created : %d record", len(toCreateConfigsDB))
	}

	Response(content, http.StatusOK, c)
}

func Response(body string, statusCode int, c *gin.Context) {
	c.String(statusCode, body)
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
