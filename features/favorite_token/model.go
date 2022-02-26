package favorite_token

type favoriteToken struct {
	WalletAddress string `json:"walletAddress" db:"wallet_address"`
	TokenAddress  string `json:"tokenAddress" db:"token_address"`
	TokenSymbol   string `json:"tokenSymbol" db:"token_symbol"`
}

var favoriteTokens []favoriteToken
