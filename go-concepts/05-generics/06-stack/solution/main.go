package main

import "fmt"

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) { s.data = append(s.data, v) }

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	last := len(s.data) - 1
	v := s.data[last]
	s.data = s.data[:last]
	return v, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

func (s *Stack[T]) Len() int { return len(s.data) }

func main() {
	si := &Stack[int]{}
	for _, n := range []int{1, 2, 3} {
		si.Push(n)
	}
	for si.Len() > 0 {
		v, _ := si.Pop()
		fmt.Println("int pop:", v)
	}

	ss := &Stack[string]{}
	ss.Push("a")
	ss.Push("b")
	if v, ok := ss.Peek(); ok {
		fmt.Println("string peek:", v) // b
	}
}
