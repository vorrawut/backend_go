package models

type FavoriteToken struct {
	WalletAddress string `json:"walletAddress" db:"wallet_address"`
	TokenAddress  string `json:"tokenAddress" db:"token_address"`
	TokenSymbol   string `json:"tokenSymbol" db:"token_symbol"`
}

var FavoriteTokens []FavoriteToken
