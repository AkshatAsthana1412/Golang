package main

import "fmt"

func main() {
	var m map[string]int

	fmt.Println("m == nil:", m == nil)
	fmt.Println("read missing key:", m["nope"]) // zero, no panic

	_, ok := m["nope"]
	fmt.Println("ok:", ok)

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("write panicked: %v\n", r)
			}
		}()
		m["x"] = 1 // panic
	}()

	m = map[string]int{}
	m["x"] = 1
	fmt.Println("after make:", m)
}
