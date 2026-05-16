// Problem 10: Request Cancellation via c.Request.Context()
//
// Every HTTP request gets a Context that is cancelled when the client
// disconnects OR the server shuts down. Long-running handlers should
// honor it: if the work is no longer needed, stop early.
//
// Tasks:
//   1. GET /work — simulate a 5-step computation, sleeping 1 second per
//      step. Between steps, check `ctx.Done()` via select; if fired,
//      return 499 (a-la nginx) with a short message.
//   2. To test: start the server, then `curl http://localhost:8080/work`
//      and Ctrl-C before 5 seconds. Server logs should show early exit.
//
// Run:
//   go run .

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	_ = r
	// TODO
}
