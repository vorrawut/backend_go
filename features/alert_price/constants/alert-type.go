package constants

type alertType struct {
	PriceRisesAbove string
	PriceDropTo     string
	ChangeIsOver    string
	ChangeIsUnder   string
}

var AlertType = alertType{
	PriceRisesAbove: "Price rises above",
	PriceDropTo:     "Price drop to",
	ChangeIsOver:    "Change is over",
	ChangeIsUnder:   "Change is under",
}
