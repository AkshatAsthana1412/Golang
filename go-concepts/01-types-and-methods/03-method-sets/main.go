// Problem 3: Method Sets
//
// The METHOD SET of a type determines which interfaces it satisfies:
//   - The method set of T includes all methods with receiver (t T).
//   - The method set of *T includes ALL methods — both (t T) and (t *T).
//
// This means a value of type T does NOT satisfy an interface that requires
// a pointer-receiver method, but a *T does.
//
// Tasks:
//   1. Define a type `Counter` with methods:
//        - Get() int            (value receiver)
//        - Increment()          (pointer receiver, mutates state)
//   2. Define an interface `Reader` { Get() int } and an interface
//      `ReadIncrementer` { Get() int; Increment() }.
//   3. Show that:
//        - `Counter{}` satisfies `Reader`
//        - `&Counter{}` satisfies BOTH `Reader` and `ReadIncrementer`
//        - `Counter{}` does NOT satisfy `ReadIncrementer` (uncomment a
//          line that should fail and observe the compile error).
//
// Why this matters:
//   - Interview classic: "why does my type not satisfy this interface?"
//   - The receiver type you pick affects how callers can use your type.
//
// Run:
//   go run .

package main

import "fmt"

// define Counter, Get, Increment.
type Counter struct{ n int }

func (c Counter) Get() int {
	return c.n
}

func (p *Counter) Increment() {
	p.n++
}

// define Reader, ReadIncrementer.
type Reader interface {
	Get() int
}

type ReadIncrementer interface {
	Get() int
	Increment()
}

func main() {
	// var r Reader = Counter{}
	var r Reader = Counter{4}
	fmt.Println(r.Get())
	// var ri ReadIncrementer = &Counter{}
	var ri ReadIncrementer = &Counter{6}
	ri.Increment()
	fmt.Printf("method Get with *Counter still works, because method set of *T includes all methods!\n%d\n", ri.Get())
	// var bad ReadIncrementer = Counter{} // should NOT compile.
	// var bad ReadIncrementer = Counter{0}
}
