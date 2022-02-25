package models

import "time"

type TelegramProfile struct {
	tableName struct{} `db:"telegram_profiles"`

	ID          int       `json:"id" db:"id"`
	TelegramID  int       `json:"telegramId" db:"telegram_id"`
	UsersID     int       `json:"usersId" db:"users_id"`
	FirstName   string    `json:"firstName" db:"first_name"`
	LastName    string    `json:"lastName" db:"last_name"`
	Username    string    `json:"userName" db:"username"`
	Hash        string    `json:"hash" db:"hash"`
	AuthDate    time.Time `json:"authDate" db:"auth_date"`
	IsActive    bool      `json:"isActive" db:"is_active"`
	CreatedDate time.Time `json:"createdDate" db:"created_date"`
	CreatedBy   string    `json:"createdBy" db:"created_by"`
	UpdatedDate time.Time `json:"updatedDate" db:"updated_date"`
	UpdatedBy   string    `json:"updatedBy" db:"updated_by"`
}
