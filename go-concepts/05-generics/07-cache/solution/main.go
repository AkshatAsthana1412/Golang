package main

import "fmt"

// Cache is a generic FIFO-eviction cache.
type Cache[K comparable, V any] struct {
	max   int
	order []K // insertion order, for FIFO eviction
	store map[K]V
}

func NewCache[K comparable, V any](max int) *Cache[K, V] {
	if max < 1 {
		max = 1
	}
	return &Cache[K, V]{
		max:   max,
		store: make(map[K]V, max),
	}
}

func (c *Cache[K, V]) Get(k K) (V, bool) {
	v, ok := c.store[k]
	return v, ok
}

func (c *Cache[K, V]) Set(k K, v V) {
	if _, exists := c.store[k]; !exists {
		if len(c.store) >= c.max {
			oldest := c.order[0]
			c.order = c.order[1:]
			delete(c.store, oldest)
		}
		c.order = append(c.order, k)
	}
	c.store[k] = v
}

func (c *Cache[K, V]) Len() int { return len(c.store) }

func main() {
	c := NewCache[string, int](3)
	c.Set("a", 1)
	c.Set("b", 2)
	c.Set("c", 3)
	c.Set("d", 4) // evicts "a"

	for _, k := range []string{"a", "b", "c", "d"} {
		v, ok := c.Get(k)
		fmt.Printf("%s -> %d (hit=%v)\n", k, v, ok)
	}
	fmt.Println("len:", c.Len())
}
