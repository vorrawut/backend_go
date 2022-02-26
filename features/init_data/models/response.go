package models

import (
	"safebsc/features/shared"
)

type InitDataResponse struct {
	SupportedTokens []shared.SupportedToken `json:"supportedTokens"`
	SupportedFarms  []shared.SupportedFarm  `json:"supportedFarms"`
}
