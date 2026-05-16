package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type ctxKey int

const reqIDKey ctxKey = 0

func slowOp(ctx context.Context, n int) (int, error) {
	if id, ok := ctx.Value(reqIDKey).(string); ok {
		fmt.Println("[slowOp] req:", id)
	}
	for i := 0; i < n; i++ {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		case <-time.After(100 * time.Millisecond):
		}
	}
	return n * 7, nil
}

func main() {
	root := context.WithValue(context.Background(), reqIDKey, "req-123")

	// Will fail: 10 * 100ms > 200ms timeout.
	c1, cancel1 := context.WithTimeout(root, 200*time.Millisecond)
	defer cancel1()
	if _, err := slowOp(c1, 10); errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("first call timed out as expected")
	}

	// Will succeed.
	c2, cancel2 := context.WithTimeout(root, 2*time.Second)
	defer cancel2()
	if v, err := slowOp(c2, 5); err == nil {
		fmt.Println("second call returned:", v)
	}
}
