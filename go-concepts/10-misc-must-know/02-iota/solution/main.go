package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	KB = 1 << (10 * (iota + 1))
	MB
	GB
	TB
)

type Perm uint8

const (
	ReadFlag Perm = 1 << iota
	WriteFlag
	ExecFlag
)

func main() {
	fmt.Println("Tuesday =", Tuesday) // 2
	fmt.Println("KB,MB,GB,TB:", KB, MB, GB, TB)

	p := ReadFlag | ExecFlag
	fmt.Printf("perm bits = %05b\n", p) // 00101
	fmt.Println("can write?", p&WriteFlag != 0)
	fmt.Println("can read? ", p&ReadFlag != 0)
}
