package models

import (
	"math"
	"math/big"
	"time"
)

type VenusPort struct {
	ProtocolName  string  `json:"protocol"`
	WalletAddress string  `json:"wallet_address"`
	Deposit       []Asset `json:"deposit"`
	Borrow        []Asset `json:"borrow"`
	UpdatedAt     float64 `json:"updated_at"`
	Reward
}

type Reward struct {
	Amount float64 `json:"reward_amount"`
	Price  float64 `json:"reward_price"`
	Value  float64 `json:"reward_value"`
}
type Asset struct {
	Address string  `json:"asset_address" bson:"asset_address"`
	Name    string  `json:"asset_name" bson:"asset_name"`
	Symbol  string  `json:"asset_symbol" bson:"asset_symbol"`
	Amount  float64 `json:"amount" bson:"asset_amount"`
	Value   float64 `json:"value" bson:"asset_value"`
	Type    string  `json:"type" bson:"type"`
	USDRate float64 `json:"rate" bson:"usd_rate"`
	APR     float64 `json:"apr"`
}

type MongoAsset struct {
	CreateAt      time.Time `json:"createAt" bson:"create_at"`
	WalletAddress string    `json:"wallet_address" bson:"wallet_addr"`
	ProtocolName  string    `bson:"protocol_name"`
	Address       string    `bson:"asset_address"`
	Name          string    `bson:"asset_name"`
	Symbol        string    `bson:"asset_symbol"`
	Amount        string    `bson:"asset_amount"`
	Value         string    `bson:"asset_value"`
	Type          string    `bson:"type"`
	USDRate       string    `bson:"usd_rate"`
	BatchNo       int       `bson:"batch_no"`
}

func (a *Asset) SetRate(rate *big.Int) {
	x := new(big.Float).SetInt(rate)
	y := new(big.Float).SetFloat64(math.Pow10(18))
	a.USDRate, _ = new(big.Float).Quo(x, y).Float64()
}
