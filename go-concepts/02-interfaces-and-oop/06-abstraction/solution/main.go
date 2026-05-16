package main

import "fmt"

type Notifier interface {
	Notify(user, message string) error
}

type EmailNotifier struct{}

func (EmailNotifier) Notify(user, message string) error {
	fmt.Printf("[email] %s: %s\n", user, message)
	return nil
}

type SMSNotifier struct{}

func (SMSNotifier) Notify(user, message string) error {
	fmt.Printf("[sms]   %s: %s\n", user, message)
	return nil
}

type OrderService struct {
	notifier Notifier // depends on the abstraction, not a concrete type
}

func NewOrderService(n Notifier) *OrderService {
	return &OrderService{notifier: n}
}

func (s *OrderService) Place(user, item string) error {
	fmt.Printf("placed: %s -> %s\n", user, item)
	return s.notifier.Notify(user, "Your order for "+item+" is confirmed.")
}

func main() {
	emailSvc := NewOrderService(EmailNotifier{})
	smsSvc := NewOrderService(SMSNotifier{})

	_ = emailSvc.Place("ada", "book")
	_ = smsSvc.Place("bob", "headphones")
}
