// Problem 6: Middleware Basics
//
// A Gin middleware is just a `gin.HandlerFunc`. Inside it you can:
//   - inspect/modify the request before the handler runs
//   - call `c.Next()` to invoke downstream handlers
//   - run code AFTER `c.Next()` returns (e.g., for timing/logging)
//   - call `c.Abort()` (or `c.AbortWithStatus`) to STOP the chain
//
// Tasks:
//   1. Write a `RequestTimer()` middleware that records `start := time.Now()`,
//      calls `c.Next()`, then logs `time.Since(start)` along with method
//      and path.
//   2. Use `r.Use(RequestTimer())` to register globally.
//   3. Add a slow handler GET /slow that sleeps 200ms; observe the
//      timer log.
//
// Run:
//   go run .
//   curl http://localhost:8080/slow

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	_ = r
	// TODO
}
