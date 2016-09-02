package spangle_org_uk

import "time"

type Question struct {
	Id int
	Text string
	Period time.Duration
	Start time.Time
	ShowLast bool
}

type Questions []Question
