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
	"time"
)

// TODO: Implement searchDB — simulates a DB search, respecting context cancellation.
// func searchDB(ctx context.Context, dbName string, query string) (string, error) { ... }

func main() {
	_ = fmt.Println
	_ = context.Background
	_ = time.Now

	// TODO:
	// 1. Create a cancellable context.
	// 2. Launch 3 goroutines searching different databases.
	// 3. Accept the first result, cancel the rest, and print the winner.
}
