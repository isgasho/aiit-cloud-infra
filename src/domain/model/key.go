package model

import "time"

type Key struct {
	ID         int
	InstanceID int
	Data       string
	CreatedAt  time.Time
}
