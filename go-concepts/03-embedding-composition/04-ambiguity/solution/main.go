package main

import "fmt"

type A struct{}

func (A) Name() string { return "A" }

type B struct{}

func (B) Name() string { return "B" }

type Outer struct {
	A
	B
}

// Without this method, `o.Name()` is ambiguous:
//   ./main.go:NN:NN: ambiguous selector o.Name
// Defining Name on Outer SHADOWS both promoted Names.
func (o Outer) Name() string { return "Outer(" + o.A.Name() + "+" + o.B.Name() + ")" }

func main() {
	o := Outer{}
	// Without the shadowing method we'd have to disambiguate:
	fmt.Println("explicit:", o.A.Name(), o.B.Name())

	// With our shadow it just works:
	fmt.Println("shadowed:", o.Name())
}
