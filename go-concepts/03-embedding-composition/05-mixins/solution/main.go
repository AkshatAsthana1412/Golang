package main

import (
	"fmt"
	"sync"
	"time"
)

type Auditable struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a *Auditable) Touch() { a.UpdatedAt = time.Now() }

type User struct {
	Auditable
	Name string
}

type Order struct {
	Auditable
	Total int
}

// Embedding *sync.Mutex by POINTER is the safe way — sync.Mutex must not
// be copied after first use, and embedding-by-value invites that bug
// (e.g., when the outer struct is passed by value somewhere).
type Cache struct {
	*sync.Mutex
	data map[string]string
}

func NewCache() *Cache {
	return &Cache{Mutex: &sync.Mutex{}, data: map[string]string{}}
}

func (c *Cache) Set(k, v string) {
	c.Lock() // promoted from *sync.Mutex
	defer c.Unlock()
	c.data[k] = v
}

func main() {
	now := time.Now()
	u := User{Auditable: Auditable{CreatedAt: now}, Name: "Ada"}
	o := Order{Auditable: Auditable{CreatedAt: now}, Total: 99}

	u.Touch() // method promotion
	o.Touch()
	fmt.Println("user updated at:", u.UpdatedAt) // field promotion
	fmt.Println("order updated at:", o.UpdatedAt)

	c := NewCache()
	c.Set("k", "v") // c.Lock()/Unlock() are promoted from the embedded Mutex
	fmt.Println("cache:", c.data)
}
