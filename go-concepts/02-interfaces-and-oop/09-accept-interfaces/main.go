// Problem 9: Accept Interfaces, Return Structs
//
// Go proverb: "Accept interfaces, return structs."
//
//   - PARAMETERS as interfaces give callers flexibility — they can pass
//     any type that satisfies the contract.
//   - RETURN values as concrete structs give callers the full API and
//     don't constrain future additions.
//
// Counter-example smell: a constructor that returns an interface forces
// every caller to use only the methods named in that interface.
//
// Tasks:
//   1. Implement `func WriteGreeting(w io.Writer, name string) error`.
//      It accepts an INTERFACE — so it works with os.Stdout, bytes.Buffer,
//      a network conn, etc.
//   2. Implement `type Counter struct { ... }` with `Inc()` / `Value()`
//      and a constructor `NewCounter() *Counter` that returns the CONCRETE
//      type. (Do not return an interface from NewCounter.)
//   3. In main(), call WriteGreeting against both os.Stdout and a
//      bytes.Buffer to demonstrate the flexibility.
//
// Run:
//   go run .

package main

import "fmt"

// TODO: WriteGreeting(io.Writer, string) error
// TODO: Counter, NewCounter

func main() {
	fmt.Println("Implement me.")
}
