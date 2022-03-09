package models

import (
	"time"
)

type WalletValueResponse struct {
	Total        string                    `json:"total" bson:"total"`
	TotalPercent string                    `json:"totalPercent" bson:"totalPercent"`
	Data         []DataWalletValueResponse `json:"data" bson:"data"`
}

type DataWalletValueResponse struct {
	Date  time.Time `json:"time" bson:"date"`
	Value string    `json:"value" bson: "value"`
}

type DataWalletValue struct {
	Date        time.Time
	SumOfValue  float64
	TotalOfData float64
	AvgValue    float64
}

type UserWallets struct {
	WalletAddress string    `json:"wallet_address" bson:"wallet_address"`
	Token         string    `json:"token" bson:"token"`
	Amount        string    `json:"amount" bson:"amount"`
	Value         string    `json:"value" bson:"value"`
	CreatedAt     time.Time `json:"createdAt" bson:"createdAt"`
}

type MyWallet struct {
	Total  int      `json:"total" bson:"total"`
	Wallet []Wallet `json:"wallet" bson:"wallet"`
}

type Wallet struct {
	Token string `json:"tokenAddress" bson:"tokenAddress"`
	Value string `json:"value" bson:"value"`
}
