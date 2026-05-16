// Problem 2: Sentinel Errors
//
// A SENTINEL ERROR is a package-level `var` that callers compare against
// to identify a specific failure mode:
//
//   var ErrNotFound = errors.New("not found")
//   if err == ErrNotFound { ... }     // older style
//   if errors.Is(err, ErrNotFound) {} // wrap-aware style (preferred)
//
// Famous examples: io.EOF, sql.ErrNoRows, os.ErrNotExist.
//
// Tasks:
//   1. Define `ErrNotFound` and `ErrPermission` as package-level errors.
//   2. Write a tiny in-memory store:
//        type Store struct { data map[string]string; allowed map[string]bool }
//        func (s *Store) Get(user, key string) (string, error)
//      It returns ErrPermission if !allowed[user], ErrNotFound if missing.
//   3. Demonstrate calling Get and matching with `errors.Is`.
//
// Run:
//   go run .

package main

import "fmt"

// TODO

func main() {
	fmt.Println("Implement me.")
}
