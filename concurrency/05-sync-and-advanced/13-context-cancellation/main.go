// Problem 13: Context Cancellation
//
// Simulate searching 3 "databases" concurrently for a record. As soon as ANY
// database returns a result, cancel the search on the other two.
//
// Setup:
//   - searchDB(ctx, dbName, query) simulates a database search with random latency.
//     It checks ctx.Done() to stop early if cancelled.
//   - Launch 3 goroutines, one per database.
//   - Use context.WithCancel to get a cancel function.
//   - When the first result arrives, call cancel() and print the result.
//
// Expected behavior:
//   - Only 1 database's result is printed.
//   - The other 2 goroutines notice cancellation and stop promptly.
//
// Hints:
//   - Create a context: ctx, cancel := context.WithCancel(context.Background())
//   - In searchDB, use a select in a loop or before the simulated work:
//       select {
//       case <-ctx.Done():
//           return "", ctx.Err()
//       case <-time.After(latency):
//           return result, nil
//       }
//   - In main, receive from the result channel once, then call cancel().
//   - defer cancel() is good practice to avoid context leaks.
//
// Run:
//   go run .

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Implement searchDB — simulates a DB search, respecting context cancellation.
func searchDB(ctx context.Context, dbName string, query string) (string, error) {
	latency := time.Duration(500+rand.Intn(500)) * time.Millisecond
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(latency):
		return fmt.Sprintf("%s returned result for query %s in %d ms.", dbName, query, latency.Milliseconds()), nil
	}
}

func main() {
	// 1. Create a cancellable context.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	results := make(chan string, 3)
	query := "select * from customers;"
	// 2. Launch 3 goroutines searching different databases.
	go func() {
		result, err := searchDB(ctx, "db1", query)
		if err == nil {
			results <- result
		}
	}()
	go func() {
		result, err := searchDB(ctx, "db2", query)
		if err == nil {
			results <- result
		}
	}()
	go func() {
		result, err := searchDB(ctx, "db3", query)
		if err == nil {
			results <- result
		}
	}()
	// 3. Accept the first result, cancel the rest, and print the winner.
	select {
	case res := <-results:
		fmt.Println(res)
		cancel()
	case <-time.After(1 * time.Second):
		fmt.Println("All searches timed out!")
	}
}
