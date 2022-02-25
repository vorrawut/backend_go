package constants

type walletType struct {
	MAIN  string
	WATCH string
}

var WalletType = walletType{
	MAIN:  "MAIN",
	WATCH: "WATCH",
}
