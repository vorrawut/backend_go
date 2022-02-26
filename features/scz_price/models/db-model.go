package models

type TokenPrice struct {
	Price  string `json:"price" bson:"price"`
	Symbol string `json:"symbol" bson:"symbol"`
	Token  string `json:"token" bson:"token"`
}
