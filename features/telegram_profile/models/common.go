package models

type CommonTelegramProfile struct {
	TelegramID int    `json:"telegramId" db:"telegram_id"`
	UsersID    int    `json:"usersId" db:"users_id"`
	FirstName  string `json:"firstName" db:"first_name"`
	LastName   string `json:"lastName" db:"last_name"`
	Username   string `json:"userName" db:"username"`
	Hash       string `json:"hash" db:"hash"`
}
