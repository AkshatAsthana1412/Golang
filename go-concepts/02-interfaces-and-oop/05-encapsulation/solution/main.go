package main

import (
	"errors"
	"fmt"
)

var (
	ErrNegativeAmount   = errors.New("amount must be positive")
	ErrInsufficientFund = errors.New("insufficient funds")
)

type BankAccount struct {
	balance int // hidden — only this package's code can read/write it
}

func NewBankAccount(initial int) (*BankAccount, error) {
	if initial < 0 {
		return nil, fmt.Errorf("initial balance: %w", ErrNegativeAmount)
	}
	return &BankAccount{balance: initial}, nil
}

func (a *BankAccount) Balance() int { return a.balance }

func (a *BankAccount) Deposit(amount int) error {
	if amount <= 0 {
		return ErrNegativeAmount
	}
	a.balance += amount
	return nil
}

func (a *BankAccount) Withdraw(amount int) error {
	if amount <= 0 {
		return ErrNegativeAmount
	}
	if amount > a.balance {
		return ErrInsufficientFund
	}
	a.balance -= amount
	return nil
}

func main() {
	a, _ := NewBankAccount(100)
	_ = a.Deposit(50)
	if err := a.Withdraw(200); err != nil {
		fmt.Printf("withdraw 200: %v\n", err)
	}
	_ = a.Withdraw(75)
	fmt.Printf("balance = %d\n", a.Balance()) // 75

	// External callers (other packages) cannot do `a.balance = -1` —
	// the field name `balance` isn't visible. That's encapsulation.
}
