// Problem 6: Abstraction via Interface
//
// Abstraction in Go = depend on BEHAVIOR (interfaces), not concrete types.
// This makes code testable and lets you swap implementations.
//
// Setup: imagine an order-service that needs to send notifications. We
// don't want it bound to "email" specifically.
//
// Tasks:
//   1. Define `Notifier` interface with `Notify(user, message string) error`.
//   2. Implement two notifiers: `EmailNotifier` and `SMSNotifier`.
//      Each just prints "[email]/[sms] to user: message".
//   3. Define `OrderService` that holds a `Notifier` (the abstraction).
//      Its `Place(user, item string) error` should print the order then
//      call `Notify`.
//   4. In main(), construct two OrderServices — one with email, one with
//      sms — and place an order with each.
//
// Why this matters:
//   - Tests can pass a fake/mock Notifier — no real email goes out.
//   - You can add a new notifier (push, slack, ...) without modifying
//     OrderService.
//
// Run:
//   go run .

package main

import "fmt"

// TODO: Notifier, EmailNotifier, SMSNotifier, OrderService
type Notifier interface {
	Notify(user, message string) error
}

type EmailNotifier struct {
}

func (e EmailNotifier) Notify(user, message string) error {
	fmt.Printf("email to user %s, message: %q\n", user, message)
	return nil
}

type SmsNotifier struct{}

func (e SmsNotifier) Notify(user, message string) error {
	fmt.Printf("sms to user %s, message: %q\n", user, message)
	return nil
}

type OrderService struct {
	Notifier
}

func NewOrderService(notifier Notifier) *OrderService {
	return &OrderService{Notifier: notifier}
}

func (o *OrderService) Place(user, message string) error {
	fmt.Println("Order created for user: ", user)
	return o.Notify(user, message)
}

func main() {
	o1 := NewOrderService(EmailNotifier{})
	o2 := NewOrderService(SmsNotifier{})
	o1.Place("Akshat", "PS5 slim")
	o2.Place("Ambrose", "Lip gloss")
}
