package model

import "time"

type State int

const (
	Starting State = iota + 1
	Initializing
	Running
	Terminating
	Terminated
)

type Instance struct {
	ID        int
	HostID    int
	Name      string
	State     State
	Size      int
	Key       *Key
	Address   *Address
	CreatedAt time.Time
}
