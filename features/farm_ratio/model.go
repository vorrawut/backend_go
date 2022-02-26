package farm_ratio

type FarmRatioResponse struct {
	Data []DataFarmRatioResponse `json:"data" bson:"data"`
}

type DataFarmRatioResponse struct {
	FarmName string `json:"farmName" bson:"farmName"`
	FarmIcon string `json:"farmIcon" bson:"farmIcon"`
	Value    string `json:"value" bson:"value"`
	Percent  string `json:"percent" bson:"percent"`
}

type DataFarmRatio struct {
	FarmId   string  `json:"farm_id" bson:"farm_id"`
	FarmName string  `json:"farmName" bson:"farmName"`
	Value    float64 `json:"avgValue" bson:"avgValue"`
}

type UserPortfolioLatests struct {
	WalletAddress string `json:"wallet_address" bson:"wallet_address"`
	FarmId        string `json:"farm_id" bson:"farm_id"`
	FarmName      string `json:"farm_name" bson:"farm_name"`
	PortValue     string `json:"port_value" bson:"port_value"`
}

type SupportedFarm struct {
	Address     string `json:"address"`
	FarmName    string `json:"farmName"`
	ChefAddress string `json:"chefAddress"`
	FarmID      string `json:"id"`
	Symbol      string `json:"tokenSymbol"`
	FarmIcon    string `json:"farmIcon"`
}
