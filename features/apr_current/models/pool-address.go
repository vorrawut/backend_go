package models

import "time"

type PoolAddress struct {
	PoolId   int32     `json:"poolId" bson:"pool_id"`
	FarmName string    `json:"farmName" bson:"farm_name"`
	Apr      float64   `json:"apr" bson:"apr"`
	CreateAt time.Time `json:"createAt" bson:"create_at"`
}
