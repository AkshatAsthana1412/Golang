// Problem 4: Custom MarshalJSON / UnmarshalJSON
//
// Implement the json.Marshaler and json.Unmarshaler interfaces to control
// exactly how a value serializes — useful for time formatting, sentinel
// values, polymorphism, and so on.
//
// Tasks:
//   1. Define `type Date struct { time.Time }`.
//   2. Implement `MarshalJSON()` to emit YYYY-MM-DD and
//      `UnmarshalJSON()` to parse the same.
//   3. Define `type Event struct { When Date `+"`"+`json:"when"`+"`"+` }` and
//      round-trip an Event through json.Marshal / json.Unmarshal.
//
// Run:
//   go run .

package main

import "fmt"

func main() {
	// TODO
	fmt.Println("Implement me.")
}
