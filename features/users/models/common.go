package models

import "time"

type Common struct {
	WalletAddress string `json:"walletAddress" db:"wallet_address"`
}

type CommonDB struct {
	Id          uint      `json:"id" db:"id"`
	CreatedDate time.Time `json:"createdDate" db:"created_date"`
	CreatedBy   string    `json:"createdBy" db:"created_by"`
	UpdatedDate time.Time `json:"updatedDate" db:"updated_date"`
	UpdatedBy   string    `json:"updatedBy" db:"updated_by"`
}
