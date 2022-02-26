package shared

import "go.mongodb.org/mongo-driver/bson/primitive"

type SupportedToken struct {
	ID             primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	OrderId        string             `json:"orderId"`
	Address        string             `json:"address"`
	Name           string             `json:"name"`
	Decimals       string             `json:"decimals"`
	Symbol         string             `json:"symbol"`
	FactoryAddress string             `json:"factoryAddress"`
	Icon           string             `json:"icon"`
	Swappable      bool               `json:"swappable"`
	IsHidden       bool               `json:"isHidden"`
	TargetToken    string             `json:"targetToken"`
}
