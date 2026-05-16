// Problem 7: Interface Composition
//
// Interfaces in Go are typically SMALL. Larger interfaces are built by
// embedding smaller ones — see io.ReadWriter:
//
//   type Reader interface { Read(p []byte) (n int, err error) }
//   type Writer interface { Write(p []byte) (n int, err error) }
//   type ReadWriter interface { Reader; Writer }
//
// Tasks:
//   1. Define small interfaces: Walker { Walk() }, Talker { Talk() string }.
//   2. Define `WalkTalker` as the composition of both.
//   3. Define a type `Robot` that has both methods. Show that *Robot or
//      Robot satisfies WalkTalker (depending on receiver choice).
//
// Run:
//   go run .

package main

import "fmt"

// TODO: Walker, Talker, WalkTalker, Robot
type Walker interface {
	Walk()
}

type Talker interface {
	Talk() string
}

type WalkTalker interface {
	Walker
	Talker
}

type Robot struct {
	name string
}

func (r Robot) Walk() {
	fmt.Printf("%s can walk\n", r.name)
}

func (r Robot) Talk() string {
	return "I can talk."
}

func main() {
	r := Robot{"Claude"}
	r.Walk()
	fmt.Printf("%q\n", r.Talk())
}
