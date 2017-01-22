package main

import (
	"fmt"
)

type State int

const (
	Running State = iota + 1
	Stopped
	Rebooting
	Terminated
)

type T struct {
	Name  string
	Port  int
	State State
}

func main() {

	state := Running

	fmt.Println("state", state)

	// not initialize state
	t := T{Name: "example", Port: 6666}
	fmt.Printf("t %+v\n", t)
}

func (s State) String() string {
	switch s {
	case Running:
		return "Running"
	case Stopped:
		return "Stopped"
	case Rebooting:
		return "Rebooting"
	case Terminated:
		return "Terminated"
	default:
		return "Unknown"
	}
}
