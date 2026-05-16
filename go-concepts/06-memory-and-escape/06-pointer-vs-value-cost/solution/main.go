package main

import (
	"fmt"
	"time"
)

type SmallStruct struct{ N int }

func (s SmallStruct) Sum() int { return s.N }

type BigStruct struct {
	A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P int64
}

func (b BigStruct) SumByValue() int64 {
	return b.A + b.B + b.C + b.D + b.E + b.F + b.G + b.H +
		b.I + b.J + b.K + b.L + b.M + b.N + b.O + b.P
}

func (b *BigStruct) SumByPointer() int64 {
	return b.A + b.B + b.C + b.D + b.E + b.F + b.G + b.H +
		b.I + b.J + b.K + b.L + b.M + b.N + b.O + b.P
}

func main() {
	const N = 5_000_000

	s := SmallStruct{N: 7}
	t := time.Now()
	var sum int
	for i := 0; i < N; i++ {
		sum += s.Sum()
	}
	fmt.Printf("Small value:    %v sum=%d\n", time.Since(t), sum)

	b := BigStruct{A: 1, P: 2}
	t = time.Now()
	var bsum int64
	for i := 0; i < N; i++ {
		bsum += b.SumByValue()
	}
	fmt.Printf("Big value:      %v\n", time.Since(t))

	t = time.Now()
	for i := 0; i < N; i++ {
		bsum += b.SumByPointer()
	}
	fmt.Printf("Big pointer:    %v\n", time.Since(t))

	// On most systems Big-by-value is a few times slower than the pointer
	// version. For Small, value receiver is a wash (or faster, depending
	// on inlining).
	_ = bsum
}
