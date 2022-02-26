package custom_token

type customToken struct {
	WalletAddress string `json:"walletAddress" db:"wallet_address"`
	TokenAddress  string `json:"tokenAddress" db:"token_address"`
	TokenSymbol   string `json:"tokenSymbol" db:"token_symbol"`
	Decimals      string `json:"decimals" db:"decimals"`
}

var customTokens []customToken
