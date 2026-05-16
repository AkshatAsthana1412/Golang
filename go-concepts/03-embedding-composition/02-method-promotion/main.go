// Problem 2: Method Promotion (and the receiver-type rule)
//
// When type T embeds type E:
//   - Methods of E with VALUE receiver  -> promoted to T  AND *T
//   - Methods of E with POINTER receiver -> promoted ONLY to *T
//
// Tasks:
//   1. Define `type Logger struct{}` with two methods:
//        - func (Logger)  Info(msg string)   // value receiver
//        - func (*Logger) Error(msg string)  // pointer receiver
//   2. Define `type Service struct { Logger }` and `Service{}.Info("ok")`
//      — works.
//   3. Try `Service{}.Error("bad")` — works ONLY if you have an addressable
//      value (Go takes the address automatically). But:
//      `Service{}` (not assigned to a variable) is NOT addressable, so
//      `Service{}.Error("bad")` fails to compile. Verify with a comment.
//   4. Build via `&Service{}` and confirm both methods work via the pointer.
//
// Run:
//   go run .

package main

import "fmt"

// TODO: Logger and methods, Service
type Logger struct{}

func (l Logger) Info(msg string) {
	fmt.Printf("%q\n", msg)
}

func (p *Logger) Error(msg string) {
	fmt.Printf("%q\n", msg)
}

type Service struct {
	Logger
}

func main() {
	Service{}.Info("Info log on value receiver")
	// Service{}.Error() -- Compile error, will take pointer type receiver, not assigned to a variable, not addressable
	svc := Service{}
	svc.Error("error") // This works because go automatically takes the address of svc
	// as the receiver
	(&Service{}).Error("This works, because now reciever is &Service{}")
}
