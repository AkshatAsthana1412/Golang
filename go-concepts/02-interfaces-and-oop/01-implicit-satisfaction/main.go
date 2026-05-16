// Problem 1: Implicit Interface Satisfaction
//
// Go has NO `implements` keyword. A type satisfies an interface if it has
// all the required methods — that's it. This is sometimes called
// "structural typing" or "duck typing at compile time".
//
// Tasks:
//   1. Define an interface `Shape { Area() float64 }`.
//   2. Define `Circle{ R float64 }` and `Square{ Side float64 }` and give
//      each an `Area()` method. Note: do NOT mention `Shape` anywhere.
//   3. Write `func TotalArea(shapes []Shape) float64` and call it with a
//      slice that mixes Circles and Squares.
//
// Why this matters:
//   - You can satisfy interfaces from packages you don't control.
//   - Tests can define narrow interfaces of just the methods THEY need.
//
// Run:
//   go run .

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	r float32
}

func (c Circle) Area() float64 {
	if c.r < 0 {
		panic("radius cannot be negative")
	}
	return float64(math.Pi * c.r * c.r)
}

type Square struct {
	a int
}

func (s Square) Area() float64 {
	if s.a < 0 {
		panic("square side cannot be negative")
	}
	return float64(s.a * s.a)
}

func TotalArea(shapes []Shape) float64 {
	var totArea float64
	for _, shape := range shapes {
		totArea += shape.Area()
	}
	return totArea
}

func main() {
	s := Square{10}
	c := Circle{5}
	fmt.Printf("Total Area: %.3f\n", TotalArea([]Shape{s, c}))
}
