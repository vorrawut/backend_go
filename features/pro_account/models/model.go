package models

import "time"

type User struct {
	TokenAddress string    `json:"tokenAddress" bson:"token_address"`
	Level        string    `json:"level" bson:"level"`
	CreatedDate  time.Time `json:"createdDate" bson:"created_date"`
	UpdatedDate  time.Time `json:"updatedDate" bson:"updated_date"`
	CreatedBy    string    `json:"createdBy" bson:"created_by"`
	UpdatedBy    string    `json:"updatedBy" bson:"updated_by"`
}

type TokenPrice struct {
	Price  string `json:"price" bson:"price"`
	Symbol string `json:"symbol" bson:"symbol"`
	Token  string `json:"token" bson:"token"`
}

type Response struct {
	Level       string `json:"level" bson:"level"`
	SczBalance  string `json:"sczBalance" bson:"scz_balance"`
	MsczBalance string `json:"msczBalance" bson:"mscz_balance"`
	Rate        string `json:"rate" bson:"rate"`
	Value       string `json:"value" bson:"value"`
}

type Users struct {
	tableName struct{} `db:"users"`

	ID            uint      `json:"id" db:"id"`
	WalletAddress string    `json:"walletAddress" db:"wallet_address"`
	Level         string    `json:"level" db:"level"`
	CreatedDate   time.Time `json:"createdDate" db:"created_date"`
	UpdatedDate   time.Time `json:"updatedDate" db:"updated_date"`
	CreatedBy     string    `json:"createdBy" db:"created_by"`
	UpdatedBy     string    `json:"updatedBy" db:"updated_by"`
}

type UsersSubscriptions struct {
	tableName struct{} `db:"users_subscriptions"`

	ID                   uint      `json:"id" db:"id"`
	UsersID              uint      `json:"usersId" db:"users_id"`
	SubscriptionsPlansID uint      `json:"subscriptionsPlansId" db:"subscriptions_plans_id"`
	StartDate            time.Time `json:"startDate" db:"start_date"`
	EndDate              time.Time `json:"endDate" db:"end_date"`
	IsActive             bool      `json:"isActive" db:"is_active"`
	CreatedDate          time.Time `json:"createdDate" db:"created_date"`
	CreatedBy            string    `json:"createdBy" db:"created_by"`
	UpdatedDate          time.Time `json:"updatedDate" db:"updated_date"`
	UpdatedBy            string    `json:"updatedBy" db:"updated_by"`
}
