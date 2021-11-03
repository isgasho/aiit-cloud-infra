package model

import "time"

type Diary struct {
	ID          int
	InstanceID  int
	Data        string
	CreatedAt   time.Time
}
