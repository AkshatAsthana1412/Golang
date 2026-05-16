// Problem 12: Graceful Shutdown
//
// `r.Run()` is convenient but doesn't give you a hook to stop accepting
// new connections while letting in-flight requests drain. The idiomatic
// pattern is to construct an `*http.Server` yourself and call
// `srv.Shutdown(ctx)` on a signal:
//
//   srv := &http.Server{Addr: ":8080", Handler: r}
//   go srv.ListenAndServe()
//   <-signals
//   ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
//   srv.Shutdown(ctx)
//
// Tasks:
//   1. Build a Gin router with a slow GET /slow handler.
//   2. Start the server in a goroutine.
//   3. Wait for SIGINT/SIGTERM via signal.NotifyContext.
//   4. On signal, call Shutdown with a 5-second timeout. Log the result.
//   5. Verify: start the server, hit /slow, then Ctrl-C before /slow finishes.
//      The handler should still complete; the server then exits.
//
// Run:
//   go run .
//   curl http://localhost:8080/slow &
//   <Ctrl-C the server>

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	_ = r
	// TODO
}
