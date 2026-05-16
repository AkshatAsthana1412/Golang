// Problem 4: Value vs Pointer Receivers
//
// Choosing the receiver type is one of the most common Go design decisions.
// Rules of thumb:
//   - Use POINTER receivers if the method MUTATES the receiver, or the
//     receiver is large and copying is expensive, or the type contains a
//     mutex (mutexes must not be copied).
//   - Use VALUE receivers for small, immutable, value-semantic types
//     (think time.Time, image.Point).
//   - Be CONSISTENT within a type — don't mix value and pointer receivers
//     unless you have a specific reason.
//
// Tasks:
//   Implement a `Wallet` type with a balance field. Provide:
//     - Balance() int             — read-only
//     - Deposit(amount int)       — mutates
//     - Withdraw(amount int) bool — mutates, returns whether it succeeded
//   Then write a tiny test in main():
//     w := Wallet{}
//     w.Deposit(100)
//     w.Withdraw(40)
//     fmt.Println(w.Balance())  // expect 60
//
// BONUS: also try writing Deposit with a VALUE receiver and observe that
// the deposit is silently lost. Note that in your comment.
//
// Run:
//   go run .

package main

import "fmt"

type Wallet struct {
	// TODO: balance field
}

// TODO: methods

func main() {
	var w Wallet
	_ = w
	fmt.Println("Implement me.")
}
