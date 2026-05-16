package main

import (
	"fmt"
	"unsafe"
)

type Bad struct {
	A bool  // 1 byte + 7 padding to align next int64
	B int64 // 8
	C bool  // 1 byte + 7 padding
	D int64 // 8
}

type Good struct {
	A bool  // 1
	C bool  // 1 (+6 padding)
	B int64 // 8
	D int64 // 8
}

func main() {
	fmt.Printf("Bad  size = %d, align = %d\n", unsafe.Sizeof(Bad{}), unsafe.Alignof(Bad{}))
	fmt.Printf("Good size = %d, align = %d\n", unsafe.Sizeof(Good{}), unsafe.Alignof(Good{}))

	fmt.Printf("bool:  size=%d align=%d\n", unsafe.Sizeof(false), unsafe.Alignof(false))
	fmt.Printf("int64: size=%d align=%d\n", unsafe.Sizeof(int64(0)), unsafe.Alignof(int64(0)))

	// Rule of thumb: order fields from largest alignment to smallest.
}
