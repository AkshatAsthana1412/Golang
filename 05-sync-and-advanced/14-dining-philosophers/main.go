// Problem 14: Dining Philosophers
//
// The classic concurrency problem:
//   - 5 philosophers sit around a table, each with a plate of food.
//   - Between each pair of philosophers is a fork (5 forks total).
//   - To eat, a philosopher needs BOTH the left and right fork.
//   - After eating, they put both forks down and think.
//   - Each philosopher eats 3 times then leaves.
//
// The naive approach (everyone picks up left fork, then right) deadlocks when
// all 5 pick up their left fork simultaneously.
//
// Your task: Implement a deadlock-free solution using asymmetric resource ordering.
// The trick: philosopher N picks up fork min(N, (N+1)%5) first, then the other.
// This breaks the circular wait condition because at least one philosopher
// picks up forks in reversed order.
//
// Hints:
//   - Represent each fork as a sync.Mutex.
//   - Each philosopher is a goroutine.
//   - Use resource ordering: always lock the lower-numbered fork first.
//   - Use a WaitGroup to wait for all philosophers to finish.
//   - Print "philosopher N is eating (round R)" and "philosopher N is thinking".
//
// Run:
//   go run .

package main

import (
	"fmt"
	"sync"
	"time"
)

// TODO: Implement the philosopher goroutine.
// func philosopher(id int, forks [5]*sync.Mutex, wg *sync.WaitGroup) { ... }

func main() {
	_ = fmt.Println
	_ = sync.Mutex{}
	_ = time.Sleep

	// TODO:
	// 1. Create 5 forks (mutexes).
	// 2. Launch 5 philosopher goroutines.
	// 3. Wait for all to finish.
}
