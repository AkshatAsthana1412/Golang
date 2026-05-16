package main

import "fmt"

func Unique[T comparable](in []T) []T {
	seen := make(map[T]struct{}, len(in))
	out := make([]T, 0, len(in))
	for _, v := range in {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	return out
}

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{m: map[T]struct{}{}}
}

func (s *Set[T]) Add(v T)       { s.m[v] = struct{}{} }
func (s *Set[T]) Has(v T) bool  { _, ok := s.m[v]; return ok }
func (s *Set[T]) Len() int      { return len(s.m) }

func main() {
	fmt.Println(Unique([]int{1, 2, 1, 3, 2, 4})) // [1 2 3 4]

	words := []string{"go", "is", "go", "is", "fast"}
	s := NewSet[string]()
	for _, w := range words {
		s.Add(w)
	}
	fmt.Println("unique words:", s.Len(), "has 'go'?", s.Has("go"))
}
