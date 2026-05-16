package main

import "fmt"

type Celsius float64    // distinct type with underlying float64
type Fahrenheit float64 // distinct type
type Kelvin = float64   // alias — Kelvin IS float64

func (c Celsius) ToF() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func main() {
	c := Celsius(100)
	var f64 float64

	// f64 = c            // ❌ compile error: cannot use Celsius as float64
	f64 = float64(c) //    ✅ explicit conversion required between distinct types
	fmt.Printf("Celsius -> float64: %v\n", f64)

	var k Kelvin = 273.15
	f64 = k //              ✅ Kelvin is just an alias for float64
	fmt.Printf("Kelvin (alias)  -> float64: %v\n", f64)

	fmt.Printf("100°C = %.2f°F\n", c.ToF())

	// You CAN'T attach methods to aliases of built-in types:
	//   func (k Kelvin) String() string { return "..." }   // ❌
	// because methods can only be defined on named types declared in the
	// same package, and Kelvin == float64 (defined in the universe block).
}
