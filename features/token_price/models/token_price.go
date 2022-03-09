package models

import (
	"time"
)

type TokenPrice struct {
	Symbol    string    `json:"symbol" bson:"symbol"`
	Price     string    `json:"price" bson:"price"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type Response struct {
	Symbol        string `json:"symbol" bson:"symbol"`
	Value         string `json:"value" bson:"value"`
	Change        string `json:"change" bson:"change"`
	PriceMovement string `json:"priceMovement" bson:"priceMovement"`
	Volume        string `json:"volume" bson:"volume"`
	Low           string `json:"low" bson:"low"`
	High          string `json:"high" bson:"high"`
}

type SupportedToken struct {
	ID             string `json:"id"`
	Address        string `json:"address"`
	Name           string `json:"name"`
	Decimals       string `json:"decimals"`
	Symbol         string `json:"symbol"`
	FactoryAddress string `json:"factoryAddress,omitempty"`
	Icon           string `json:"icon"`
	IsHidden       bool   `json:"isHidden,omitempty"`
}
