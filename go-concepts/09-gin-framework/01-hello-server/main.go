// Problem 1: Hello Gin Server
//
// Gin is a lightweight HTTP framework built on net/http with a fast
// radix-tree router and ergonomic helpers (c.JSON, c.Bind*, middleware).
//
// Tasks:
//   1. Create a `gin.Default()` engine (which comes with Logger + Recovery
//      middleware pre-installed).
//   2. Add a GET /ping route that returns `{"message":"pong"}` with HTTP 200.
//   3. Add a GET /healthz route returning 204 No Content (no body).
//   4. r.Run(":8080").
//
// Verify:
//   curl -i http://localhost:8080/ping
//   curl -i http://localhost:8080/healthz
//
// Run:
//   go run .

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	_ = r

	// TODO: routes
	// TODO: r.Run(":8080")
}
