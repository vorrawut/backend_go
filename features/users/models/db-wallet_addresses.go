package models

type WalletAddresses struct {
	tableName struct{} `db:"wallet_addresses"`

	Common
	CommonDB
}
