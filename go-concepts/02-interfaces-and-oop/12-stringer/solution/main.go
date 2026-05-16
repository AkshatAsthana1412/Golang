package main

import "fmt"

type Status int

const (
	Active Status = iota
	Inactive
	Banned
)

func (s Status) String() string {
	switch s {
	case Active:
		return "[active]"
	case Inactive:
		return "[inactive]"
	case Banned:
		return "[banned]"
	default:
		return "[?]"
	}
}

func main() {
	for _, s := range []Status{Active, Inactive, Banned, Status(99)} {
		fmt.Printf("%%v = %v   %%d = %d\n", s, s)
	}

	// Note: %d ignores Stringer and formats the int directly. This is
	// occasionally useful (debugging the raw enum value).
}
