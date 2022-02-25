package models

type UsersWalletAddresses struct {
	tableName struct{} `db:"users_wallet_addresses"`

	Common
	UsersId           uint   `json:"usersId" db:"users_id"`
	WalletAddressesId uint   `json:"walletAddressesId" db:"wallet_addresses_id"`
	WalletType        string `json:"walletType" db:"wallet_type"`
	CommonDB
}
