package shared

import "go.mongodb.org/mongo-driver/bson/primitive"

type SupportedFarm struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	OrderID     string             `json:"orderId"`
	FarmName    string             `json:"farmName"`
	FarmIcon    string             `json:"farmIcon"`
	ChefAddress string             `json:"chefAddress"`
	ABI         []struct {
		Inputs []struct {
			InternalType string `json:"internalType"`
			Name         string `json:"name"`
			Type         string `json:"type"`
		} `json:"inputs"`
		StateMutability string        `json:"stateMutability,omitempty"`
		Type            string        `json:"type"`
		Anonymous       bool          `json:"anonymous,omitempty"`
		Name            string        `json:"name,omitempty"`
		Outputs         []interface{} `json:"outputs,omitempty"`
	} `json:"ABI"`
	TokenName     string `json:"tokenName"`
	TokenAddress  string `json:"tokenAddress"`
	TokenSymbol   string `json:"tokenSymbol"`
	TokenDecimals int    `json:"tokenDecimals"`
	TokenImage    string `json:"tokenImage"`
	PendingStmt   string `json:"pendingStmt"`
	Withdraw      string `json:"withdraw"`
	Harvest       string `json:"harvest"`
	Sos           string `json:"sos"`
	Website       string `json:"website"`
	Factory       string `json:"factory"`
	FactoryABI    []struct {
		Inputs []struct {
			InternalType string `json:"internalType"`
			Name         string `json:"name"`
			Type         string `json:"type"`
		} `json:"inputs"`
		StateMutability string `json:"stateMutability,omitempty"`
		Type            string `json:"type"`
		Anonymous       bool   `json:"anonymous,omitempty"`
		Name            string `json:"name,omitempty"`
		Outputs         []struct {
			InternalType string `json:"internalType"`
			Name         string `json:"name"`
			Type         string `json:"type"`
		} `json:"outputs,omitempty"`
	} `json:"factoryABI"`
	Pool0Withdraw string `json:"pool0Withdraw"`
	PriceContract bool   `json:"priceContract"`
	Class         string `json:"class"`
}
