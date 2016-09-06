package questions

import "time"

type Question struct {
	Id		int		`json:"id"`
	Text		string		`json:"text"`
	Period		time.Duration	`json:"period"`
	Start		time.Time	`json:"start"`
	ShowLast	bool		`json:"show_last"`
}

type Questions []Question
