// Problem 5: Encapsulation via Unexported Fields
//
// Go's access control is at the PACKAGE level:
//   - Identifiers starting with an uppercase letter are EXPORTED.
//   - Identifiers starting with a lowercase letter are PACKAGE-PRIVATE.
//
// Encapsulation = expose behavior, hide state. The classic pattern is:
//   - Unexported struct fields
//   - Constructor function `NewThing(...)` that enforces invariants
//   - Exported methods that mediate access
//
// Tasks:
//   Build a `BankAccount` with:
//     - unexported field `balance int`
//     - `NewBankAccount(initial int)` that rejects negative starting balance
//     - `Balance() int`, `Deposit(amount int) error`, `Withdraw(amount int) error`
//     - errors must reject non-positive amounts and overdrafts
//
// Run:
//   go run .

package main

import (
	"errors"
	"fmt"
)

var (
	ErrorInvalidAmount       = errors.New("invalid amount, amount must be positive")
	ErrorInsufficientBalance = errors.New("insufficient balance")
)

// TODO: BankAccount struct, NewBankAccount, methods
type BankAccount struct {
	balance int
}

func validateAmount(amt int) error {
	if amt <= 0 {
		return ErrorInvalidAmount
	}
	return nil
}

func NewBankAccount(initial int) (*BankAccount, error) {
	if err := validateAmount(initial); err != nil {
		return &BankAccount{}, err
	}
	return &BankAccount{balance: initial}, nil
}

func (b *BankAccount) Deposit(amt int) error {
	if err := validateAmount(amt); err != nil {
		return err
	}
	b.balance += amt
	fmt.Println("Deposit successful! New Balance: ", b.balance)
	return nil
}

func (b *BankAccount) Withdraw(amt int) error {
	if err := validateAmount(amt); err != nil {
		return err
	}
	if amt > b.balance {
		return ErrorInsufficientBalance
	}
	b.balance -= amt
	fmt.Println("Withdraw successful! Remaining balance: ", b.balance)
	return nil
}

func main() {
	_, err := NewBankAccount(-19)
	if err != nil {
		fmt.Println("Error creating bank account: ", err)
	}
	b, err := NewBankAccount(1000)
	if err != nil {
		fmt.Println("Error creating bank account: ", err)
	}
	err = b.Deposit(100)
	if err != nil {
		fmt.Println("Error depositing: ", err)
	}
	err = b.Withdraw(199)
	if err != nil {
		fmt.Println("Error withdrawing: ", err)
	}
	err = b.Withdraw(-100)
	if err != nil {
		fmt.Println("Error withdrawing: ", err)
	}
	err = b.Deposit(234)
	if err != nil {
		fmt.Println("Error depositing: ", err)
	}
	err = b.Withdraw(23343)
	if err != nil {
		fmt.Println("Error withdrawing: ", err)
	}
}
