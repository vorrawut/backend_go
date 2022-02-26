package reward_value

import (
	"time"
)

type InvestmentReportResponse struct {
	Total        string                         `json:"total"`
	TotalPercent string                         `json:"totalPercent"`
	Data         []DataInvestmentReportResponse `json:"data" bson:"data"`
}

type DataInvestmentReportResponse struct {
	Date   time.Time `json:"time"`
	Reward string    `json:"value"`
}

type DataInvestmentReport struct {
	Date   time.Time
	Reward float64
}

type UserPortfolios struct {
	WalletAddress string    `json:"wallet_address" bson:"wallet_address"`
	PortAmount    string    `json:"port_amount" bson:"port_amount"`
	PortValue     string    `json:"port_value" bson:"port_value"`
	FarmId        string    `json:"farm_id" bson:"farm_id"`
	FarmName      string    `json:"farm_name" bson:"farm_name"`
	PoolId        float32   `json:"pool_id" bson:"pool_id"`
	LpAddress     string    `json:"lp_address" bson:"lp_address"`
	Reward        []Reward  `json:"reward" bson:"reward"`
	CreatedAt     time.Time `json:"createdAt" bson:"createdAt"`
}

type UserPortfolioLatests struct {
	WalletAddress string   `json:"wallet_address" bson:"wallet_address"`
	Reward        []Reward `json:"reward" bson:"reward"`
}

type Reward struct {
	Amount string `json:"amount" bson:"amount"`
	Value  string `json:"value" bson:"value"`
}
