// Problem 9: Centralized Error Handling
//
// Calling `c.Error(err)` adds `err` to a per-request error list (visible
// via `c.Errors`). A global middleware can inspect those errors AFTER
// `c.Next()` and produce a consistent error envelope, removing the need
// for each handler to format its own error responses.
//
// Tasks:
//   1. Define `var ErrNotFound = errors.New("not found")` and a custom
//      `*BadInput{Field string}` error.
//   2. Write `ErrorMiddleware()` that runs after handlers and, if
//      `len(c.Errors) > 0`, picks the LAST error and maps to:
//        ErrNotFound  -> 404
//        *BadInput    -> 400 with {"field": ...}
//        else         -> 500
//      Make sure to use `c.JSON` only if no response was written yet
//      (`!c.Writer.Written()`).
//   3. /missing handler: c.Error(ErrNotFound)
//      /invalid handler: c.Error(&BadInput{Field: "name"})
//      /boom handler:    c.Error(errors.New("kaboom"))
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
