package model

import "time"

type Host struct {
	ID        int
	Name      string
	Limit     int
	CreatedAt time.Time
}
