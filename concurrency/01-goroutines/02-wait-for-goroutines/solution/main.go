// Solution 2: Wait For Goroutines
//
// Key takeaways:
//   - sync.WaitGroup is a counting semaphore: Add(n) increments the counter,
//     Done() decrements it, Wait() blocks until the counter reaches zero.
//   - Always call Add() BEFORE launching the goroutine (not inside it) to avoid
//     a race where Wait() sees zero before the goroutine has called Add().
//   - `defer wg.Done()` ensures Done is called even if the goroutine panics.
//   - All 5 fetches run concurrently, so total wall-clock time ≈ slowest single
//     fetch rather than the sum of all fetches.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

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

	var wg sync.WaitGroup

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
