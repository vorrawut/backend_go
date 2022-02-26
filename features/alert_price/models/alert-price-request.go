package models

type AlertPriceRequest struct {
	CommonAlertPrice
	AlertType int `json:"alertType"`
	Frequency int `json:"frequency"`
}
