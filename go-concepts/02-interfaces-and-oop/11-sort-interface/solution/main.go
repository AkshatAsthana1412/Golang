package main

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	people := []Person{
		{"Cara", 30}, {"Ada", 27}, {"Bob", 41}, {"Don", 27},
	}

	// Style 1: sort.Interface
	sort.Sort(ByAge(people))
	fmt.Println("by age (sort.Interface):", people)

	// Style 2: sort.Slice (closure on the slice)
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("by name (sort.Slice):  ", people)

	// Style 3: generic slices.SortFunc with cmp.Compare
	slices.SortFunc(people, func(a, b Person) int {
		// Sort by age desc, breaking ties on name asc.
		if c := cmp.Compare(b.Age, a.Age); c != 0 {
			return c
		}
		return cmp.Compare(a.Name, b.Name)
	})
	fmt.Println("by age desc, name (slices.SortFunc):", people)
}
