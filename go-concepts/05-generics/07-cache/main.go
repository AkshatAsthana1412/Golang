// Problem 7: Generic Capped Cache
//
// Tasks:
//   Implement `Cache[K comparable, V any]` with:
//     - NewCache[K, V](max int) *Cache[K, V]
//     - Get(k K) (V, bool)
//     - Set(k K, v V)            // when full, evict an arbitrary key
//     - Len() int
//   For "evict an arbitrary key", just delete the first key the map yields
//   on iteration — Go intentionally randomizes map iteration order, which
//   is fine for this exercise.
//
//   Optional: keep a slice of keys in insertion order for FIFO eviction.
//
// Run:
//   go run .

package main

import "fmt"

// TODO

func main() {
	fmt.Println("Implement me.")
}
