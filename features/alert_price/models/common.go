package models

type CommonAlertPrice struct {
	TokenAddress string  `json:"tokenAddress" db:"token_address"`
	Symbol       string  `json:"symbol" db:"symbol"`
	Value        float64 `json:"value" db:"value"`
	UsersID      uint    `json:"usersId" db:"users_id"`
}
