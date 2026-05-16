package main

import "fmt"

type Point struct{ X, Y int }

type Tags struct {
	values []string
}

type Holder struct {
	Any interface{}
}

type sliceFoo struct{ s []int }

func (sliceFoo) Foo() {}

func main() {
	// Plain comparable struct.
	fmt.Println("Point eq:", Point{1, 2} == Point{1, 2})

	// Tags{} == Tags{}             // ❌ compile: invalid op: struct containing []string cannot be compared
	// We can use reflect.DeepEqual or slices.Equal instead — see Topic 07.

	// Interface fields make comparability a RUNTIME affair.
	a := Holder{Any: sliceFoo{s: []int{1}}}
	b := Holder{Any: sliceFoo{s: []int{1}}}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Holder == panicked: %v\n", r)
		}
	}()

	// Compiles fine — but the dynamic type sliceFoo contains a slice, so
	// the interface comparison panics: "comparing uncomparable type ...".
	fmt.Println("Holder eq:", a == b)
}
