package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var i int
	var f float64
	var b bool
	var s string
	var p *int
	var sl []int
	var m map[string]int
	var ch chan int
	var err error
	var pers Person

	fmt.Printf("int       = %v\n", i)
	fmt.Printf("float64   = %v\n", f)
	fmt.Printf("bool      = %v\n", b)
	fmt.Printf("string    = %q\n", s)
	fmt.Printf("*int      = %v (nil? %v)\n", p, p == nil)
	fmt.Printf("[]int     = %v (nil? %v, len=%d)\n", sl, sl == nil, len(sl))
	fmt.Printf("map       = %v (nil? %v)\n", m, m == nil)
	fmt.Printf("chan      = %v (nil? %v)\n", ch, ch == nil)
	fmt.Printf("error     = %v (nil? %v)\n", err, err == nil)
	fmt.Printf("struct    = %+v\n", pers)

	// Nil slice append is fine — append allocates a new backing array.
	sl = append(sl, 1, 2, 3)
	fmt.Printf("appended nil slice -> %v\n", sl)

	// Nil map write panics.
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("nil map write panicked: %v\n", r)
			}
		}()
		m["x"] = 1
	}()

	// Reading from a nil map returns the zero value (no panic).
	v := m["missing"]
	fmt.Printf("nil map read returned zero value: %v\n", v)
}
