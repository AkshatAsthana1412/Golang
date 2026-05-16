package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct{ R float64 }

func (c Circle) Area() float64 { return math.Pi * c.R * c.R }

type Square struct{ Side float64 }

func (s Square) Area() float64 { return s.Side * s.Side }

func TotalArea(shapes []Shape) float64 {
	var total float64
	for _, sh := range shapes {
		total += sh.Area()
	}
	return total
}

func main() {
	shapes := []Shape{
		Circle{R: 2},
		Square{Side: 3},
		Circle{R: 1},
	}
	fmt.Printf("total = %.3f\n", TotalArea(shapes))

	// Notice neither Circle nor Square mentions Shape — yet both satisfy it.
}
