package token_price_graph

import (
	"time"
)

type tokenPrice struct {
	Symbol    string    `json:"symbol" bson:"symbol"`
	Price     string    `json:"price" bson:"price"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type response struct {
	Time  time.Time `json:"time" bson:"time"`
	Low   string    `json:"low" bson:"low"`
	High  string    `json:"high" bson:"high"`
	Open  string    `json:"open" bson:"open"`
	Close string    `json:"close" bson:"close"`
}
