package entity

import "time"

type Report struct {
	ID         string
	CreatedAt  time.Time
	FinishedAt time.Time
	Status     string
}
