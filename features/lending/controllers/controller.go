package controllers

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"math/big"
	"net/http"
	"safebsc/mongo"
	"safebsc/utils"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	mongodriver "go.mongodb.org/mongo-driver/mongo"

	"safebsc/features/lending/models"
	std "safebsc/features/lending/std-reference"
	venus "safebsc/features/lending/venus-comptroller"
)

var globalClient *ethclient.Client

const DISTRIBUTION = "0xfd36e2c2a6789db23113685031d7f16329158384"
const STD_REFERENCE = "0xDA7a001b254CD22e46d3eAB04d937489c93174C3"

type Controller struct{}

func (u Controller) GetVenus(c *gin.Context) {
	getVenus(c)
}

func init() {
	var err error
	globalClient, err = ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		panic(err)
	}
}

func getVenus(c *gin.Context) {
	address := c.Query("address")
	if !utils.IsValidAddress(address) {
		c.String(http.StatusBadRequest, "Invalid wallet address.")
	}

	rwch := make(chan models.Reward)
	go getReward(address, rwch)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoDB, err := mongo.Connect()
	if err != nil {
		log.Fatal(err)
		Response(err.Error(), 500, c)
	}
	defer mongoDB.Client.Disconnect(mongoDB.Ctx)

	latest := findLatestBatch(ctx, mongoDB.Client.Database(mongoDB.DatabaseName), address)
	log.Println("now :", time.Now().Unix(), "latest record :", latest)
	if latest < 0 {
		Response("[]", 200, c)
	}

	if float64(time.Now().Unix())-latest/1000 > 3600 {
		Response("[]", 200, c)
	}

	mongoAssets := findAssets(ctx, mongoDB.Client.Database(mongoDB.DatabaseName), address, latest)
	venusPort := transformMongoAssetToVenus(mongoAssets, latest)
	venusPort.Reward = <-rwch
	content, err := json.Marshal(venusPort)
	if err != nil {
		log.Fatal(err)
		Response(err.Error(), 500, c)
	}

	Response(string(content), 200, c)
}

func transformMongoAssetToVenus(mgas []models.MongoAsset, timestamp float64) (vp models.VenusPort) {
	vp.ProtocolName = "Venus Protocol"
	vp.WalletAddress = mgas[0].WalletAddress
	vp.UpdatedAt = timestamp
	for _, mga := range mgas {
		if value, err := strconv.ParseFloat(mga.Value, 64); err != nil || value <= 0 {
			continue
		}
		amount, _ := strconv.ParseFloat(mga.Amount, 64)
		value, _ := strconv.ParseFloat(mga.Value, 64)
		usdRate, _ := strconv.ParseFloat(mga.USDRate, 64)

		if mga.Type == "Deposit" {
			vp.Deposit = append(vp.Deposit, models.Asset{
				Address: mga.Address,
				Name:    mga.Name,
				Symbol:  mga.Symbol,
				Amount:  amount,
				Value:   value,
				Type:    mga.Type,
				USDRate: usdRate,
				APR:     0,
			})
		}
		if mga.Type == "Borrow" {
			vp.Borrow = append(vp.Borrow, models.Asset{
				Address: mga.Address,
				Name:    mga.Name,
				Symbol:  mga.Symbol,
				Amount:  amount,
				Value:   value,
				Type:    mga.Type,
				USDRate: usdRate,
				APR:     0,
			})
		}
	}

	return vp
}

func findAssets(ctx context.Context, db *mongodriver.Database, address string, latest float64) []models.MongoAsset {
	lendingPortCollection := db.Collection("user_lending_portfolios")
	filterCursor, err := lendingPortCollection.Find(ctx, bson.M{"wallet_addr": address, "batch_no": latest})
	if err != nil {
		log.Fatal(err)
	}
	var mongoAssets = make([]models.MongoAsset, 0)
	if err = filterCursor.All(ctx, &mongoAssets); err != nil {
		log.Fatal(err)
	}
	return mongoAssets
}

func findLatestBatch(ctx context.Context, db *mongodriver.Database, address string) float64 {
	lendingPortCollection := db.Collection("user_lending_portfolios")
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"wallet_addr": address,
			},
		},
		{
			"$group": bson.M{
				"_id": "",
				"max": bson.M{"$max": "$batch_no"}},
		},
	}
	maxCursor, err := lendingPortCollection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err)
	}
	result := []bson.M{}
	if err = maxCursor.All(ctx, &result); err != nil {
		log.Fatal(err)
	}

	if len(result) == 0 {
		return -1
	}

	latest := result[0]["max"].(float64)
	return latest
}

func getReward(address string, rwch chan models.Reward) {
	var wg sync.WaitGroup
	defer close(rwch)

	var rate float64
	var amount float64

	wg.Add(1)
	go func() {
		stdClient, err := std.NewStdreference(common.HexToAddress(STD_REFERENCE), globalClient)
		if err != nil {
			log.Println(err)
		}
		price, err := stdClient.GetReferenceData(&bind.CallOpts{}, "XVS", "USD")
		if err != nil {
			log.Println(err)
		}
		x := new(big.Float).SetInt(price.Rate)
		y := new(big.Float).SetFloat64(math.Pow10(18))
		rate, _ = new(big.Float).Quo(x, y).Float64()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		venusClient, err := venus.NewVenuscomptrollerCaller(common.HexToAddress(DISTRIBUTION), globalClient)
		if err != nil {
			log.Println(err)
		}
		a, err := venusClient.VenusAccrued(&bind.CallOpts{}, common.HexToAddress(address))
		if err != nil {
			log.Println(err)
		}
		x := new(big.Float).SetInt(a)
		y := new(big.Float).SetFloat64(math.Pow10(18))
		amount, _ = new(big.Float).Quo(x, y).Float64()
		wg.Done()
	}()

	wg.Wait()

	rwch <- models.Reward{
		Amount: amount,
		Price:  rate,
		Value:  amount * rate,
	}
}

func Handler(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		getVenus(c)

	default:
		c.String(http.StatusMethodNotAllowed, "Error")
	}
}

func Response(body string, statusCode int, c *gin.Context) {
	c.String(statusCode, body)
}
