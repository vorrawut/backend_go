package models

import "time"

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

type SubscriptionsPlans struct {
	tableName struct{} `db:"subscriptions_plans"`

	ID                  uint      `json:"id" db:"id"`
	Description         string    `json:"description" db:"description"`
	Amount              float64   `json:"amount" bson:"amount"`
	SmartContractPlanID float64   `json:"smartContractPlanId" bson:"smart_contract_plan_id"`
	MerchantAddress     string    `json:"merchantAddress" db:"merchant_address"`
	Frequency           float64   `json:"frequency" db:"frequency"`
	TargetToken         string    `json:"targetToken" db:"target_token"`
	CreatedDate         time.Time `json:"createdDate" db:"created_date"`
	CreatedBy           string    `json:"createdBy" db:"created_by"`
	UpdatedDate         time.Time `json:"updatedDate" db:"updated_date"`
	UpdatedBy           string    `json:"updatedBy" db:"updated_by"`
}

type UsersSubscriptions struct {
	tableName struct{} `db:"users_subscriptions"`

	ID          uint      `json:"id" db:"id"`
	IsActive    bool      `json:"isActive" db:"is_active"`
	CreatedDate time.Time `json:"createdDate" db:"created_date"`
	CreatedBy   string    `json:"createdBy" db:"created_by"`
	UpdatedDate time.Time `json:"updatedDate" db:"updated_date"`
	UpdatedBy   string    `json:"updatedBy" db:"updated_by"`
	RegisterSubscriptionRequest
}

type RegisterSubscriptionRequest struct {
	UsersID              uint      `json:"usersId" db:"users_id"`
	SubscriptionsPlansID uint      `json:"subscriptionsPlansId" db:"subscriptions_plans_id"`
	StartDate            time.Time `json:"startDate" db:"start_date"`
	EndDate              time.Time `json:"endDate" db:"end_date"`
}
