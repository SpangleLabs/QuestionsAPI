package questions

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	Text		string		`json:"text"`
	Period		time.Duration	`json:"period"`
	Start		time.Time	`json:"start"`
	ShowLast	bool		`json:"show_last"`
}

type Questions []Question
