package constants

type frequencyType struct {
	OnlyOnce      string
	OnceADay      string
	Always        string
	ChangeIsUnder string
}

var FrequencyType = frequencyType{
	OnlyOnce: "Only Once",
	OnceADay: "Once a day",
	Always:   "Always",
}
