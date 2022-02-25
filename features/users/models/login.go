package models

type LoginRequest struct {
	Common
}

type LoginResponse struct {
	UsersId    uint `json:"usersId"`
	TelegramID int  `json:"telegramId"`
}
