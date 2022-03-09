package models

import (
	"time"
)

type FarmValueResponse struct {
	Total        string                  `json:"total"`
	TotalPercent string                  `json:"totalPercent"`
	Data         []DataFarmValueResponse `json:"data"`
}

type DataFarmValueResponse struct {
	Date  time.Time `json:"time"`
	Value string    `json:"value"`
}

type DataFarmValue struct {
	Date  time.Time
	Value float64
	Farms []Farm
}

type Farm struct {
	FarmId      string
	SumOfValue  float64
	TotalOfData float64
	AvgValue    float64
}

type UserPortfolios struct {
	WalletAddress string    `json:"wallet_address" bson:"wallet_address"`
	PortAmount    string    `json:"port_amount" bson:"port_amount"`
	PortValue     string    `json:"port_value" bson:"port_value"`
	FarmId        string    `json:"farm_id" bson:"farm_id"`
	FarmName      string    `json:"farm_name" bson:"farm_name"`
	PoolId        float32   `json:"pool_id" bson:"pool_id"`
	LpAddress     string    `json:"lp_address" bson:"lp_address"`
	CreatedAt     time.Time `json:"createdAt" bson:"createdAt"`
}

type UserPortfolioLatests struct {
	WalletAddress string `json:"wallet_address" bson:"wallet_address"`
	PortValue     string `json:"port_value" bson:"port_value"`
}
