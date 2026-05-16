// Problem 6: Panic vs Error — Which to Use When
//
// Idiomatic Go:
//   - Use `error` returns for EXPECTED failure modes (network down, bad
//     input, missing key, ...). Callers expect to handle them.
//   - Use `panic` only for INVARIANT VIOLATIONS — bugs ("this can't happen"),
//     fatal config issues at startup, or unrecoverable corruption.
//   - The standard library uses panic in contexts like `regexp.MustCompile`
//     where misuse is a programmer error, not a runtime condition.
//
// Tasks:
//   1. Implement `MustEnv(key string) string` that calls `os.Getenv(key)`
//      and PANICS if empty — patterned after MustCompile.
//   2. Implement `Env(key string) (string, error)` that returns an error
//      instead. Show how each is intended to be used.
//   3. Discuss in a comment: when would a library author choose Must*
//      vs the error-returning version?
//
// Run:
//   PORT=8080 go run .
//   (then: go run .  to see it panic)

package main

import "fmt"

// TODO: MustEnv, Env

func main() {
	fmt.Println("Implement me.")
}
