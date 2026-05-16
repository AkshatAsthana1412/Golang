package main

import "fmt"

func Double(s []int) {
	for i := range s {
		s[i] *= 2
	}
}

func AppendOne(s []int) {
	s = append(s, 99)
	_ = s
}

func AppendOneRet(s []int) []int {
	return append(s, 99)
}

func main() {
	a := []int{1, 2, 3}
	Double(a)
	fmt.Println("after Double:", a) // [2 4 6]

	AppendOne(a)
	fmt.Println("after AppendOne:", a) // unchanged length

	a = AppendOneRet(a)
	fmt.Println("after AppendOneRet:", a)
}
