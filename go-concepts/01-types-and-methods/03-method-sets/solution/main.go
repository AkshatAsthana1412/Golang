package main

import "fmt"

type Counter struct{ n int }

func (c Counter) Get() int  { return c.n } // value receiver
func (c *Counter) Increment() { c.n++ }    // pointer receiver

type Reader interface{ Get() int }
type ReadIncrementer interface {
	Get() int
	Increment()
}

func main() {
	// Method set of Counter (value): {Get}
	var r Reader = Counter{}
	fmt.Println("Counter{} as Reader:", r.Get())

	// Method set of *Counter: {Get, Increment}
	c := &Counter{}
	c.Increment()
	c.Increment()
	var ri ReadIncrementer = c
	fmt.Println("*Counter as ReadIncrementer:", ri.Get())

	// Compile error if uncommented:
	//   cannot use Counter{} as ReadIncrementer: missing method Increment
	//     (Increment has pointer receiver).
	// var bad ReadIncrementer = Counter{}
	// _ = bad

	// Note: Go DOES auto-address when calling pointer-receiver methods on
	// addressable VALUES — that's why `someCounter.Increment()` works on
	// a local variable. But interface assignment is NOT a method call,
	// so the addressability shortcut doesn't apply.
}
