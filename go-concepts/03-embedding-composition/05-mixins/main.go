// Problem 5: Mixins via Embedding
//
// Embedding lets you add cross-cutting capabilities to a type without
// repeating boilerplate. Common examples in Go:
//   - sync.Mutex embedded in a type that needs internal locking
//   - sync.WaitGroup embedded similarly
//   - A "Base" struct holding common audit fields (CreatedAt, UpdatedAt)
//
// Tasks:
//   1. Define `Auditable { CreatedAt, UpdatedAt time.Time }` with a method
//      `Touch()` that sets UpdatedAt to time.Now().
//   2. Define `User { Auditable; Name string }`.
//   3. Define `Order { Auditable; Total int }`.
//   4. In main(), create one of each, call .Touch() (promoted), and print
//      .UpdatedAt (also promoted).
//   5. Bonus: embed a `*sync.Mutex` (pointer) into a `Cache` type and
//      observe that you can call `cache.Lock()` directly. Why pointer
//      rather than value? Because a sync.Mutex must NOT be copied.
//
// Run:
//   go run .

package main

import "fmt"

// TODO: Auditable, User, Order, Cache (with embedded *sync.Mutex)

func main() {
	fmt.Println("Implement me.")
}
