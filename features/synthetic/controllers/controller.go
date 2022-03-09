package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"safebsc/features/postgres"
	"safebsc/features/synthetic/chainlink"
	"safebsc/features/synthetic/twindex"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Base struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type SyntheticToken struct {
	Base
	TokenName   string  `json:"tokenName" `
	TokenSymbol string  `json:"tokenSymbol"`
	TokenAddr   string  `json:"tokenAddress"`
	ChainAddr   string  `json:"chainAddress"`
	ChainPrice  float64 `json:"chainPrice"`
	PoolPrice   float64 `json:"poolPrice"`
	Premium     string  `json:"premium"`
	Icon        string  `json:"icon"`
}

type tokenConfig struct {
	tokenName string
	tokenAddr string
	chainAddr string
	icon      string
}

var clts = []tokenConfig{
	{
		tokenName: "AAPL",
		tokenAddr: "0xc10b2ce6a2bcfdfdc8100ba1602c1689997299d3",
		chainAddr: "0xb7Ed5bE7977d61E83534230f3256C021e0fae0B6",
		icon:      "/assets/images/daapl.png",
	}, {
		tokenName: "TSLA",
		tokenAddr: "0x17ace02e5c8814bf2ee9eaaff7902d52c15fb0f4",
		chainAddr: "0xEEA2ae9c074E87596A85ABE698B2Afebc9B57893",
		icon:      "/assets/images/dtsla.png",
	}, {
		tokenName: "GOOGL",
		tokenAddr: "0x9c169647471c1c6a72773cffc50f6ba285684803",
		chainAddr: "0xeDA73F8acb669274B15A977Cb0cdA57a84F18c2a",
		icon:      "/assets/images/dgoogl.png",
	}, {
		tokenName: "AMZN",
		tokenAddr: "0x1085b90544ff5c421d528aaf79cc65afc920ac79",
		chainAddr: "0x51d08ca89d3e8c12535BA8AEd33cDf2557ab5b2a",
		icon:      "/assets/images/damzn.png",
	}}

var globalClient *ethclient.Client

type Controller struct{}

func (u Controller) GetSynthetic(c *gin.Context) {
	getSynthetic(c)
}

func init() {
	var err error
	globalClient, err = ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		panic(err)
	}
}

func getSynthetic(c *gin.Context) {
	db := postgres.Connect()
	defer db.Close()

	sttDB := make([]SyntheticToken, 0)

	readFromDB := false
	// TODO : set readFromDB to true and uncomment to read data from DB
	// db.AutoMigrate(&SyntheticToken{})
	// db.Model(&SyntheticToken{}).Find(&sttDB)

	if len(sttDB) > 0 && readFromDB {
		content, err := json.Marshal(sttDB)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}

		c.String(http.StatusOK, string(content))
	}

	syntChan := make(chan SyntheticToken)

	go calls(clts, syntChan)

	stt := make([]SyntheticToken, 0)
	for synt := range syntChan {
		stt = append(stt, synt)
	}

	content, err := json.Marshal(stt)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	c.String(http.StatusOK, string(content))
}

func calls(s []tokenConfig, syntChan chan SyntheticToken) {
	var wg sync.WaitGroup
	for _, v := range s {
		wg.Add(1)
		go call(v, syntChan, &wg)
	}
	defer func() {

		wg.Wait()
		close(syntChan)
	}()
}

func call(v tokenConfig, c chan SyntheticToken, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	chainPriceChan := make(chan float64)
	go func(chainAddress string) {
		contractAddress := common.HexToAddress(chainAddress)
		instance, err := chainlink.NewChainlink(contractAddress, globalClient)
		if err != nil {
			log.Fatal(err)
		}
		lrd, err := instance.LatestRoundData(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}
		chainPriceChan <- float64(lrd.Answer.Int64()) / 100_000_000

	}(v.chainAddr)

	poolPriceChan := make(chan float64)
	go func(tokenAddress string) {
		twindexContractAddress := common.HexToAddress("0x6B011d0d53b0Da6ace2a3F436Fd197A4E35f47EF")
		twinInstance, err := twindex.NewTwindex(twindexContractAddress, globalClient)
		if err != nil {
			log.Fatal(err)
		}
		dollyAddress := common.HexToAddress("0xff54da7caf3bc3d34664891fc8f3c9b6dea6c7a5")
		syntAddress := common.HexToAddress(tokenAddress)
		amounts, err := twinInstance.GetAmountsOut(&bind.CallOpts{}, big.NewInt(10_000), []common.Address{syntAddress, dollyAddress})
		if err != nil {
			log.Fatal(err)
		}
		poolPriceChan <- float64(amounts[1].Int64()) / 10_000
	}(v.tokenAddr)

	chainPrice := <-chainPriceChan
	poolPrice := <-poolPriceChan

	premuim := (poolPrice - chainPrice) / chainPrice * 100
	c <- SyntheticToken{
		TokenName:   v.tokenName,
		TokenSymbol: v.tokenName,
		TokenAddr:   v.tokenAddr,
		ChainAddr:   v.chainAddr,
		ChainPrice:  chainPrice,
		PoolPrice:   poolPrice,
		Premium:     fmt.Sprintf("%f", premuim),
		Icon:        v.icon,
	}
}

func Handler(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		getSynthetic(c)

	default:
		c.String(http.StatusMethodNotAllowed, "Error")
	}
}
