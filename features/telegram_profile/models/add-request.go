package models

type AddTelegramProfileRequest struct {
	CommonTelegramProfile
	AuthDate string `json:"authDate" db:"auth_date"`
}
