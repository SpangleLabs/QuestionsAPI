package spangle_org_uk

import "time"

type Answer struct {
	Id		int		`json:"id"`
	Date		time.Time	`json:"date_time"`
	Question	Question	`json:"question"`
	Text		string		`json:"text"`
}

type Answers []Answer
