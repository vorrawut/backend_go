package models

import "time"

type AlertPrice struct {
	tableName struct{} `db:"alert_prices"`

	CommonAlertPrice
	ID            int       `json:"id" db:"id"`
	TargetPrice   float64   `json:"targetPrice" db:"target_price"`
	AlertType     string    `json:"alertType" db:"alert_type"`
	Frequency     string    `json:"frequency" db:"frequency"`
	AlertDateTime time.Time `json:"alertDateTime" db:"alert_date_time"`
	IsActive      bool      `json:"isActive" db:"is_active"`
	CreatedDate   time.Time `json:"createdDate" db:"created_date"`
	CreatedBy     string    `json:"createdBy" db:"created_by"`
	UpdatedDate   time.Time `json:"updatedDate" db:"updated_date"`
	UpdatedBy     string    `json:"updatedBy" db:"updated_by"`
}
