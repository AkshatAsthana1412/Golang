package main

import "fmt"

type Point struct{ X, Y int }

// localPoint: Point is built and returned by VALUE. The caller's stack
// receives a copy. No escape needed.
func localPoint() Point {
	return Point{X: 1, Y: 2}
}

// escapingPoint: returns a POINTER to a local. The local must outlive
// the function -> escapes to heap. With -gcflags="-m" you'll see:
//   "moved to heap: p" or "&Point{...} escapes to heap".
func escapingPoint() *Point {
	p := Point{X: 3, Y: 4}
	return &p
}

// viaInterface: storing into an `any` typically causes the value to escape
// because the interface header carries a pointer to the underlying value
// and the compiler usually can't prove it stays on-frame.
func viaInterface(p any) {
	_ = p
}

func main() {
	a := localPoint()
	b := escapingPoint()
	viaInterface(a)
	fmt.Println(a, *b)

	// Try:
	//   go build -gcflags="-m=2" ./solution
	// to see verbose escape decisions.
}
