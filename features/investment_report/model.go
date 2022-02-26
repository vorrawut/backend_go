package investment_report

import (
	"time"
)

type InvestmentReportResponse struct {
	Data []DataInvestmentReportResponse `json:"data" bson:"data"`
}

type DataInvestmentReportResponse struct {
	NetWorth string    `json:"netWorth" bson:"netWorth"`
	Reward   string    `json:"reward" bson: "reward"`
	Date     time.Time `json:"date" bson:"date"`
}

type DataInvestmentReport struct {
	Date     time.Time
	NetWorth float64
	Reward   float64
	Farms    []Farm
}

type Farm struct {
	FarmId      string
	NetWorth    float64
	Reward      float64
	TotalOfData float64
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

type Reward struct {
	Amount string `json:"amount" bson:"amount"`
	Value  string `json:"value" bson:"value"`
}
