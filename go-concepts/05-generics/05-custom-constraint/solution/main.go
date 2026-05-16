package main

import "fmt"

type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Sum[T Number](xs []T) T {
	var s T
	for _, v := range xs {
		s += v
	}
	return s
}

type Celsius float64

func main() {
	fmt.Println(Sum([]int{1, 2, 3}))             // 6
	fmt.Println(Sum([]float64{1.5, 2.5}))         // 4
	fmt.Printf("%T %v\n", Sum([]Celsius{20, 22}), Sum([]Celsius{20, 22}))
	// main.Celsius 42  — return type follows the input element type.
}
