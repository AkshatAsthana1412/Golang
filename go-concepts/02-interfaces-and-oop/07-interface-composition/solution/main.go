package main

import "fmt"

type Walker interface {
	Walk()
}

type Talker interface {
	Talk() string
}

// WalkTalker is the composition.
type WalkTalker interface {
	Walker
	Talker
}

type Robot struct{ Name string }

func (r Robot) Walk()         { fmt.Printf("%s walks.\n", r.Name) }
func (r Robot) Talk() string  { return r.Name + " says hello" }

func main() {
	var wt WalkTalker = Robot{Name: "R2"}
	wt.Walk()
	fmt.Println(wt.Talk())

	// We can also satisfy just Walker if a context only needs walking:
	var w Walker = Robot{Name: "C3"}
	w.Walk()
}
