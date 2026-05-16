package main

import "fmt"

type Logger struct{}

func (Logger) Info(msg string)   { fmt.Println("INFO ", msg) }
func (*Logger) Error(msg string) { fmt.Println("ERROR", msg) }

type Service struct{ Logger }

func main() {
	// Value receiver method — works on any Service.
	Service{}.Info("starting") // not even addressable; value receiver doesn't care

	// Pointer receiver method on an ADDRESSABLE value (variable):
	s := Service{}
	s.Error("oops") // Go takes &s automatically because s is addressable

	// Same method on a NON-addressable composite literal:
	// Service{}.Error("nope") // ❌ compile: cannot take address of Service{}

	// Always works via an explicit pointer:
	(&Service{}).Error("via ptr")
}
