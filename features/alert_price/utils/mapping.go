package utils

import "safebsc/features/alert_price/constants"

func MapFrequency(frequency int) string {
	result := ""
	switch frequency {
	case 1:
		result = constants.FrequencyType.OnlyOnce
		break
	case 2:
		result = constants.FrequencyType.OnceADay
		break
	case 3:
		result = constants.FrequencyType.Always
		break
	default:
		result = ""
		break
	}
	return result
}

func MapFrequencyByName(frequency string) int {
	result := 0
	switch frequency {
	case constants.FrequencyType.OnlyOnce:
		result = 1
		break
	case constants.FrequencyType.OnceADay:
		result = 2
		break
	case constants.FrequencyType.Always:
		result = 3
		break
	default:
		result = 0
		break
	}
	return result
}

func MapAlertType(alertType int) string {
	result := ""
	switch alertType {
	case 1:
		result = constants.AlertType.PriceRisesAbove
		break
	case 2:
		result = constants.AlertType.PriceDropTo
		break
	case 3:
		result = constants.AlertType.ChangeIsOver
		break
	case 4:
		result = constants.AlertType.ChangeIsUnder
		break
	default:
		result = ""
		break
	}
	return result
}

func MapAlertTypeByName(alertType string) int {
	result := 0
	switch alertType {
	case constants.AlertType.PriceRisesAbove:
		result = 1
		break
	case constants.AlertType.PriceDropTo:
		result = 2
		break
	case constants.AlertType.ChangeIsOver:
		result = 3
		break
	case constants.AlertType.ChangeIsUnder:
		result = 4
		break
	default:
		result = 0
		break
	}
	return result
}
