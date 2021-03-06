package models

import (
	"safebsc/features/shared"
	"time"
)

type UserPortFolio struct {
	BatchNo       int64     `json:"batch_no" bson:"batch_no"`
	WalletAddress string    `json:"wallet_address" bson:"wallet_address"`
	PortValue     string    `json:"port_value" bson:"port_value"`
	FarmID        string    `json:"farm_id" bson:"farm_id"`
	CreatedAt     time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt" bson:"updatedAt"`
	Reward        []Reward  `json:"reward" bson:"reward"`
	PoolID        int       `json:"pool_id" bson:"pool_id"`
	LpTokens      []LpToken `json:"lp_tokens" bson:"lp_tokens"`
	LpAddress     string    `json:"lp_address" bson:"lp_address"`
	PortAmount    string    `json:"portAmount" bson:"port_amount"`
}

type Reward struct {
	Symbol  string `json:"symbol" bson:"symbol"`
	Address string `json:"address" bson:"address"`
	Amount  string `json:"amount" bson:"amount"`
	Value   string `json:"value" bson:"value"`
}

type LpToken struct {
	Symbol string `json:"symbol" `
	Amount string `json:"amount" `
	Value  string `json:"value"`
}

type MyFarm struct {
	FarmID            string               `json:"farmId" bson:"id"`
	Name              string               `json:"name" bson:"name"`
	ImageURL          string               `json:"imageUrl" `
	TotalPrice        string               `json:"totalPrice" `
	WalletAddress     string               `json:"walletAddress"`
	FarmDetails       []FarmDetail         `json:"farmDetails" `
	FarmInfo          shared.SupportedFarm `json:"farmInfo"`
	UpdatedAt         time.Time            `json:"updatedAt"`
	TokenPrice        string               `json:"tokenPrice"`
	DailyRewardAmount string               `json:"dailyRewardAmount"`
}

type Token struct {
	Name          string  `json:"name"`
	Symbol        string  `json:"symbol"`
	ImageURL      string  `json:"imageUrl"`
	Balance       float64 `json:"-"`
	BalanceUsd    float64 `json:"-"`
	BalanceStr    string  `json:"balance"`
	BalanceUsdStr string  `json:"balanceUsd"`
	RawBalance    string  `json:"rawBalance"`
}

type FarmDetail struct {
	Name              string           `json:"name"`
	PoolID            int              `json:"poolId"`
	Reward            []RewardResponse `json:"reward"`
	Tokens            []Token          `json:"tokens"`
	LpAddress         string           `json:"lpAddress" `
	BalanceUsd        float64          `json:"-"`
	BalanceUsdStr     string           `json:"balanceUsd" `
	BalanceStr        string           `json:"balance" `
	Apr               string           `json:"apr"`
	AprDaily          string           `json:"aprDaily"`
	DailyRewardAmount string           `json:"dailyRewardAmount"`
	PortAmount        string           `json:"portAmount"`
}

type RewardResponse struct {
	Symbol       string  `json:"symbol"`
	Reward       float64 `json:"-"`
	RewardUsd    float64 `json:"-"`
	RewardStr    string  `json:"reward"`
	RewardUsdStr string  `json:"rewardUsd"`
}

type SupportedFarm struct {
	Address     string `json:"address"`
	FarmName    string `json:"farmName"`
	ChefAddress string `json:"chefAddress"`
	FarmID      string `json:"id"`
	Symbol      string `json:"tokenSymbol"`
	ImageURL    string `json:"farmIcon"`
}

type TokenPrice struct {
	Price  string `json:"price" bson:"price"`
	Symbol string `json:"symbol" bson:"symbol"`
	Token  string `json:"token" bson:"token"`
}

type FarmPoolApr struct {
	PoolId      int    `json:"poolId" bson:"poolId"`
	PoolAddress string `json:"pool_address" bson:"pool_address"`
	Apr         string `json:"apr" bson:"apr"`
}

type FarmApr struct {
	FarmName  string      `json:"farm_name" bson:"farm_name"`
	UpdatedAt time.Time   `json:"updatedAt" bson:"updatedAt"`
	Pool      FarmPoolApr `json:"pool" bson:"pool"`
}
