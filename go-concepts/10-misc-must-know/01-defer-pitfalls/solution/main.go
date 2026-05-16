package main

import "fmt"

func argsCapturedAtDeferSite() {
	fmt.Println("\n(a) defer args evaluated at defer site:")
	for i := 0; i < 4; i++ {
		defer fmt.Println("  i =", i) // i is EVALUATED now
	}
	// Prints (LIFO): i=3, i=2, i=1, i=0
}

func closureCapturesByReference() {
	fmt.Println("\n(b) defer closure captures by reference:")
	for i := 0; i < 4; i++ {
		defer func() { fmt.Println("  closure i =", i) }()
	}
	// In Go 1.22+: per-iteration variable. Prints 3,2,1,0.
	// Pre-1.22:    shared variable. Would print 4,4,4,4.
}

func loopFileLeakPattern() {
	fmt.Println("\n(c) loop-defer leak shape:")
	files := []string{"a", "b", "c"}
	for _, name := range files {
		// Bad: every Close runs at FUNCTION exit, not loop iteration end.
		// for _, name := range files {
		//   f, _ := os.Open(name)
		//   defer f.Close()       // accumulates, all run at end
		// }

		// Good: wrap the work in a closure so defer runs per-iteration.
		func(n string) {
			// open(n), defer close, do work, close runs HERE
			fmt.Println("  processing", n)
		}(name)
	}
}

func main() {
	argsCapturedAtDeferSite()
	closureCapturesByReference()
	loopFileLeakPattern()
}
