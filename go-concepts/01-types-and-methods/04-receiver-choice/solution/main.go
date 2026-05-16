package main

import "fmt"

type Wallet struct {
	balance int // unexported — only Wallet's methods can mutate it
}

func (w Wallet) Balance() int { return w.balance } // value receiver: read-only

func (w *Wallet) Deposit(amount int) { // pointer receiver: mutates
	w.balance += amount
}

func (w *Wallet) Withdraw(amount int) bool {
	if amount > w.balance {
		return false
	}
	w.balance -= amount
	return true
}

// Anti-example — value receiver "deposit" mutates a COPY, lost on return.
func (w Wallet) DepositBuggy(amount int) {
	w.balance += amount // modifies the copy; original is untouched
}

func main() {
	w := Wallet{}
	w.Deposit(100)
	ok := w.Withdraw(40)
	fmt.Printf("withdraw ok=%v balance=%d\n", ok, w.Balance()) // 60

	w.DepositBuggy(1000) // silently lost
	fmt.Printf("after buggy deposit, balance=%d\n", w.Balance()) // still 60

	ok = w.Withdraw(9999)
	fmt.Printf("over-withdraw ok=%v balance=%d\n", ok, w.Balance())
}
