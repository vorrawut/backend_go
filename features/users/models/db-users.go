package models

type Users struct {
	tableName struct{} `db:"users"`

	Common
	TelegramAccount string `json:"telegramAccount" db:"telegram_account"`
	LineToken       string `json:"lineToken" db:"line_token"`
	Level           string `json:"level" db:"level"`
	CommonDB
}
