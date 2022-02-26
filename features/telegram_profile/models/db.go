package models

import "time"

type TelegramProfile struct {
	tableName struct{} `db:"telegram_profiles"`

	CommonTelegramProfile
	ID          int       `json:"id" db:"id"`
	AuthDate    time.Time `json:"authDate" db:"auth_date"`
	IsActive    bool      `json:"isActive" db:"is_active"`
	CreatedDate time.Time `json:"createdDate" db:"created_date"`
	CreatedBy   string    `json:"createdBy" db:"created_by"`
	UpdatedDate time.Time `json:"updatedDate" db:"updated_date"`
	UpdatedBy   string    `json:"updatedBy" db:"updated_by"`
}
