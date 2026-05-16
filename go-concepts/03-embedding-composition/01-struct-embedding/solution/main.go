package main

import "fmt"

type Engine struct{ HP int }

func (e Engine) Start() { fmt.Printf("vroom (%d hp)\n", e.HP) }

type Car struct {
	Engine // embedded — gives Car the HP field and Start() method
	Wheels int
}

func DriveCar(c Car) { c.Start(); fmt.Printf("...on %d wheels\n", c.Wheels) }

func main() {
	c := Car{Engine: Engine{HP: 200}, Wheels: 4}

	// Field promotion:
	fmt.Println("c.HP        =", c.HP)
	fmt.Println("c.Engine.HP =", c.Engine.HP) // same value, different path

	// Method promotion:
	c.Start() // == c.Engine.Start()

	DriveCar(c)

	// DriveCar(Engine{HP: 200}) // ❌ compile: an Engine is not a Car.
	// Embedding gives composition, not inheritance — Car *contains* Engine.
}
