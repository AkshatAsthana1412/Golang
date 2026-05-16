package main

import "fmt"

func printAny(x any) {
	fmt.Println(x) // x is already boxed
}

func printT[T any](x T) {
	fmt.Println(x)
}

func main() {
	for i := 0; i < 3; i++ {
		printAny(i) // each call typically allocates: int -> interface{}
	}
	for i := 0; i < 3; i++ {
		printT(i) // generics avoid the box (depending on shape stencils)
	}

	// Why this matters in real code:
	//   logger.Info("count", count)   // any-typed args -> alloc per call
	// Hot paths often switch to typed loggers (e.g., zerolog/slog with
	// typed attribute funcs) precisely to avoid boxing.
}
