// Problem 2: Type Definitions vs Type Aliases
//
// Go has TWO syntaxes that look similar but mean very different things:
//
//   type Celsius   float64    // type DEFINITION — new distinct type
//   type Fahrenheit = float64 // type ALIAS      — same type, new name
//
// Tasks:
//   1. Define `Celsius` as a new type with underlying `float64`.
//   2. Define `Kelvin` as an alias for `float64`.
//   3. Try to assign a `Celsius` value to a `float64` variable directly
//      (no conversion). Does it compile? Add a comment with the answer.
//   4. Try to assign a `Kelvin` value to a `float64` variable directly.
//      Does it compile? Add a comment.
//   5. Add a method `(c Celsius) ToF() Fahrenheit` (treat Fahrenheit as
//      its own defined type). Note that you cannot define methods on
//      aliases of built-in types — only on named types from the same package.
//
// Why this matters:
//   - Type definitions create distinct types you can attach methods to and
//     get type-safety benefits from (e.g., can't pass meters where seconds
//     are expected).
//   - Aliases are mostly used for gradual API migration / refactoring.
//
// Run:
//   go run .

package main

import "fmt"

// TODO: type Celsius ...
// TODO: type Fahrenheit ...
// TODO: type Kelvin = float64

// TODO: func (c Celsius) ToF() Fahrenheit { ... }

func main() {
	// TODO: demonstrate the conversion rules above.
	fmt.Println("Implement me.")
}
