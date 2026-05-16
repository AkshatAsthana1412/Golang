// Problem 12: fmt.Stringer & Custom Formatting
//
// Implementing `fmt.Stringer` makes your type print nicely:
//
//   type Stringer interface { String() string }
//
// Whenever fmt prints your value with %v / %s / Println, it calls String().
//
// Bonus: implementing `fmt.Formatter` (a richer interface) lets you handle
// %v, %d, %#v, etc., separately. We won't go that deep here.
//
// Tasks:
//   1. Define `type Status int` and constants Active, Inactive, Banned
//      with iota.
//   2. Implement `String()` on Status. Print `[active]` / `[inactive]` /
//      `[banned]` for the three values, `[?]` otherwise.
//   3. Print a Status with both `%v` and `%d` and observe that %d skips
//      String() — it formats the underlying int.
//
// Run:
//   go run .

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
	for _, s := range []Status{Active, Inactive, Banned, Status(33)} {
		fmt.Printf("%%v = %v   %%d = %d\n", s, s)
	}
}
