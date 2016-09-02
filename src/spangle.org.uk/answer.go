package spangle_org_uk

import "time"

type Answer struct {
	Id int
	Date time.Time
	Question Question
	Text string
}

type Answers []Answer
