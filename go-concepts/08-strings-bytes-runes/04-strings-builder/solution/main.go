package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var b strings.Builder
	b.Grow(1000 * 5) // approximate

	for i := 1; i <= 1000; i++ {
		if i > 1 {
			b.WriteByte(',')
		}
		b.WriteString(strconv.Itoa(i))
	}

	out := b.String()
	fmt.Printf("len=%d head=%q tail=%q\n", len(out), out[:20], out[len(out)-15:])

	// Don't do this:
	//   b2 := b      // vet: assignment copies lock value to b2
	//   _ = b2.String()
	// strings.Builder uses a noCopy sentinel because String() relies on
	// the buffer pointer staying unique.
}
