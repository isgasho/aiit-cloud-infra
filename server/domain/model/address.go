package model

import "time"

type Address struct {
	ID         int
	IPAddress  string
	MacAddress string
	InstanceID int
	CreatedAt  time.Time
}
