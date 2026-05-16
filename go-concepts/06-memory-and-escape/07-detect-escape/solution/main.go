package main

import "fmt"

type Big struct{ buf [256]int }

// Returning a pointer to a local — definitely escapes.
func makeBig() *Big {
	b := Big{}
	for i := range b.buf {
		b.buf[i] = i
	}
	return &b
}

// Returning by value — caller's frame holds the data; no escape required.
// (Yes the value is large; that's a different cost — copy vs alloc.)
func makeBigByValue() Big {
	b := Big{}
	for i := range b.buf {
		b.buf[i] = i
	}
	return b
}

func main() {
	p := makeBig()
	v := makeBigByValue()
	fmt.Println(p.buf[100], v.buf[100])

	// Sample output of `go build -gcflags="-m" ./solution`:
	//   ./main.go:NN:6: can inline makeBig
	//   ./main.go:NN:9: moved to heap: b
	//   ./main.go:NN:6: can inline makeBigByValue
	//   ./main.go:NN:6: makeBigByValue ignoring self-assignment ...
	//
	// "moved to heap: b" is the headline you're looking for.
}
