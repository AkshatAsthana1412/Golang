// Problem 2: Wait For Goroutines
//
// You have a list of URLs to "fetch" (simulated with time.Sleep).
// Fetch all of them concurrently and print the result of each fetch.
// Use sync.WaitGroup to wait for all goroutines to finish before main exits.
//
// Hints:
//   - Create a WaitGroup: var wg sync.WaitGroup
//   - Before launching each goroutine, call wg.Add(1)
//   - Inside each goroutine, defer wg.Done()
//   - After launching all goroutines, call wg.Wait() to block until they're done
//   - The simulateFetch function is provided — just call it from each goroutine
//
// Run:
//   go run .

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// simulateFetch pretends to fetch a URL and returns a status message.
func simulateFetch(url string) string {
	delay := time.Duration(100+rand.Intn(400)) * time.Millisecond
	time.Sleep(delay)
	return fmt.Sprintf("fetched %s in %v", url, delay)
}

func main() {
	urls := []string{
		"https://example.com/api/users",
		"https://example.com/api/posts",
		"https://example.com/api/comments",
		"https://example.com/api/albums",
		"https://example.com/api/photos",
	}

	wg := sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := simulateFetch(url)
			fmt.Println(result)
		}()
	}

	wg.Wait()
	fmt.Println("All fetches complete.")
}
