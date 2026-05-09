// Solution 13: Context Cancellation
//
// Key takeaways:
//   - context.WithCancel returns a derived context and a cancel function. When
//     cancel() is called, ctx.Done() is closed — all goroutines watching it
//     can detect this and stop work.
//   - This is cooperative cancellation: goroutines must explicitly check
//     ctx.Done() (via select). If they don't check, cancel has no effect.
//   - Always defer cancel() to prevent context leaks, even if you call it
//     explicitly earlier.
//   - The pattern generalizes: context.WithTimeout and context.WithDeadline
//     cancel automatically after a duration/at a time. All three are used
//     heavily in production Go code (HTTP handlers, gRPC, database queries).

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type dbResult struct {
	db     string
	record string
	err    error
}

func searchDB(ctx context.Context, dbName string, query string) (string, error) {
	latency := time.Duration(100+rand.Intn(900)) * time.Millisecond

	select {
	case <-ctx.Done():
		fmt.Printf("  %s: cancelled before completing\n", dbName)
		return "", ctx.Err()
	case <-time.After(latency):
		return fmt.Sprintf("[%s] found '%s' in %v", dbName, query, latency), nil
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	databases := []string{"postgres", "redis", "elasticsearch"}
	query := "user:42"

	resultCh := make(chan dbResult, len(databases))

	for _, db := range databases {
		go func() {
			record, err := searchDB(ctx, db, query)
			resultCh <- dbResult{db: db, record: record, err: err}
		}()
	}

	// Accept the first successful result.
	first := <-resultCh
	cancel() // signal other goroutines to stop

	if first.err != nil {
		fmt.Println("error:", first.err)
	} else {
		fmt.Println("winner:", first.record)
	}

	// Give cancelled goroutines a moment to print their cancellation messages.
	time.Sleep(100 * time.Millisecond)
}
