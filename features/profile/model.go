package profile

import "time"

type user struct {
	TokenAddress    string    `json:"tokenAddress" bson:"token_address"`
	IsPremium       bool      `json:"isPremium" bson:"is_premium"`
	TelegramAccount string    `json:"telegramAccount" bson:"telegram_account"`
	LineToken       string    `json:"lineToken" bson:"line_token"`
	Level           string    `json:"level" bson:"level"`
	Price           string    `json:"price" bson:"price"`
	ProgramStart    time.Time `json:"programStart" bson:"program_start"`
	ProgramEnd      time.Time `json:"programEnd" bson:"program_end"`
	CreatedDate     time.Time `json:"createdDate" bson:"created_date"`
	UpdatedDate     time.Time `json:"updatedDate" bson:"updated_date"`
	CreatedBy       string    `json:"createdBy" bson:"created_by"`
	UpdatedBy       string    `json:"updatedBy" bson:"updated_by"`
}
