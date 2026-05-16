// Problem 7: Build Tags / Conditional Compilation
//
// You can include or exclude files from a build using:
//
//   //go:build linux
//   //go:build linux && amd64
//   //go:build dev
//   //go:build integration
//
// Place the directive at the TOP of the file, before the package clause,
// followed by a BLANK LINE. Tags can be passed via -tags:
//
//   go build -tags integration .
//
// Tasks:
//   1. This problem is split into 3 files:
//        - main.go         (always built)
//        - feature_dev.go  (//go:build dev)
//        - feature_prod.go (//go:build !dev)
//      Each "feature" file defines `func featureName() string`.
//   2. main calls featureName(). Build with and without `-tags dev`:
//        go run .
//        go run -tags dev .
//      and observe the different output.
//
// Run:
//   go run .
//   go run -tags dev .

package main

import "fmt"

func main() {
	fmt.Println("active feature:", featureName())
}
