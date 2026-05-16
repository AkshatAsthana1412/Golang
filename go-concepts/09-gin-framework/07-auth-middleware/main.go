// Problem 7: Auth Middleware + c.Set / c.Get
//
// Middleware can stash values into the request context with `c.Set(key, val)`
// and downstream handlers retrieve them with `c.Get(key)` (or
// `c.MustGet`). This is the standard pattern for "authenticate once,
// then handlers see the user."
//
// Tasks:
//   1. Write `AuthRequired()` that reads `Authorization: Bearer <token>`
//      and:
//        - if missing/invalid: 401 + abort
//        - else `c.Set("userID", <id>)` and Next()
//      Stub validation: any non-empty token after "Bearer " is treated
//      as `userID = "user-" + token`.
//   2. Apply only to /api/*. Public endpoint /public must work without
//      auth.
//   3. /api/me returns the current userID via `c.MustGet("userID")`.
//
// Verify:
//   curl http://localhost:8080/public
//   curl -i http://localhost:8080/api/me
//   curl -i http://localhost:8080/api/me -H 'Authorization: Bearer abc'
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
