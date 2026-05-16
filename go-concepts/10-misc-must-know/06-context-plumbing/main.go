// Problem 6: context.Context Plumbing
//
// `context.Context` is the standard way to pass:
//   - cancellation signals
//   - deadlines / timeouts
//   - request-scoped values (sparingly!)
//
// Convention: pass ctx as the FIRST argument of any function that does
// I/O or could be cancelled.
//
// Tasks:
//   1. `func slowOp(ctx context.Context, n int) (int, error)` does
//      `time.Sleep(n*100ms)` but checks ctx.Done() in the meantime —
//      return ctx.Err() if cancelled.
//   2. In main, call slowOp twice:
//        - with `context.WithTimeout(ctx, 200ms)` and n=10 -> should fail
//        - with `context.WithTimeout(ctx, 2s)` and n=5 -> should succeed
//   3. Show context-value plumbing: stash a request-id with
//      `context.WithValue(ctx, key, "req-123")`, retrieve in slowOp.
//      Use a CUSTOM unexported KEY TYPE to avoid collisions.
//
// Run:
//   go run .

package main

import "fmt"

func main() {
	fmt.Println("Implement me.")
}
