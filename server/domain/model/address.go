package model

import "time"

type Address struct {
	ID         int
	HostID     int
	IPAddress  string
	MacAddress string
	InstanceID int
	CreatedAt  time.Time
}
