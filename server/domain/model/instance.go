package model

import (
	"fmt"
	"time"
)

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

func StringToState(s string) (State, error) {
	switch s {
	case Starting.String():
		return Starting, nil
	case Initializing.String():
		return Initializing, nil
	case Running.String():
		return Running, nil
	case Terminating.String():
		return Terminating, nil
	case Terminated.String():
		return Terminated, nil
	}
	return 0, fmt.Errorf("%v is invalid State", s)
}
