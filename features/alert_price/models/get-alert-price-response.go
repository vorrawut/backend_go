package models

type GetAlertPriceResponse struct {
	Symbol          string              `json:"symbol"`
	TokenAddress    string              `json:"tokenAddress"`
	Price           string              `json:"price"`
	AlertPriceItems []GetAlertPriceItem `json:"alertPriceItems"`
}

type GetAlertPriceItem struct {
	ID          string `json:"id"`
	Value       string `json:"value"`
	TargetPrice string `json:"targetPrice"`
	AlertType   string `json:"alertType"`
	Frequency   string `json:"frequency"`
}
