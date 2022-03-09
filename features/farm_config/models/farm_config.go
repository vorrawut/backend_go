package models

import "time"

type Base struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type FarmConfig struct {
	Base
	Address        string `json:"address" bson:"address" binding:"required" `
	FarmName       string `json:"farmName" bson:"farm_name" binding:"required"`
	TokenName      string `json:"tokenName" bson:"token_name"`
	FactoryAddress string `json:"factoryAddress" bson:"factory_address"`
	FarmID         string `json:"farmId" bson:"farm_id" binding:"required"`
	ChefAddress    string `json:"chefAddress" bson:"chef_address"`
	IsDisable      *bool  `json:"isDisable" bson:"is_disable" binding:"required" pg:",use_zero"`
}
