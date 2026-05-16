package main

import "fmt"

var leak []*[]byte // package-level escape sink

// (a) Returning the slice itself — it ESCAPES because the backing array
// must outlive the function.
func makeBufA() []byte {
	return make([]byte, 64)
}

// (b) Closure captures the local. The closure is returned, so the local
// must live as long as the closure does -> escape.
func makeBufB() func() []byte {
	buf := make([]byte, 64)
	return func() []byte { return buf }
}

// (c) Stash into a package-level slice — definitely escapes.
func makeBufC() {
	buf := make([]byte, 64)
	leak = append(leak, &buf)
}

func main() {
	a := makeBufA()
	b := makeBufB()()
	makeBufC()
	fmt.Printf("len(a)=%d len(b)=%d len(leak)=%d\n", len(a), len(b), len(leak))
}
