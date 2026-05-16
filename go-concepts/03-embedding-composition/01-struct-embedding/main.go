// Problem 1: Struct Embedding
//
// Go has no inheritance. Instead it has EMBEDDING: declaring a field by
// type name only (without a field name) gives you composition with
// automatic method/field promotion.
//
//   type Engine struct { HP int }
//   func (e Engine) Start() { ... }
//
//   type Car struct {
//     Engine    // embedded
//     Wheels int
//   }
//
//   c := Car{Engine: Engine{HP: 200}, Wheels: 4}
//   c.HP        // promoted: same as c.Engine.HP
//   c.Start()   // promoted: same as c.Engine.Start()
//
// Tasks:
//   1. Define `Engine` and `Car` as above.
//   2. Show field access via both `c.HP` (promoted) and `c.Engine.HP` (full).
//   3. Show method call via promotion.
//   4. Embedding is COMPOSITION, not inheritance — verify that a func
//      `func DriveCar(c Car)` cannot be called with an Engine alone.
//
// Run:
//   go run .

package main

import "fmt"

// TODO: Engine, Car
type Engine struct {
	HP int
}

func (e Engine) Start() {
	fmt.Println("Engine starting!")
}

type Car struct {
	Engine
	Wheels int
}

func DriveCar(c Car) {
	fmt.Println("Driving car@!")
}

func main() {
	c := Car{Engine: Engine{HP: 900}, Wheels: 4}
	fmt.Println("Car has horsepower: ", c.HP)
	c.Engine.Start()
	DriveCar(c)
}
