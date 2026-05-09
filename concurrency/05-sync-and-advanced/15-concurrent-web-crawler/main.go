// Problem 15: Concurrent Web Crawler
//
// Crawl a simulated web graph concurrently with:
//   - Bounded parallelism (at most N concurrent fetches)
//   - URL deduplication (never visit the same URL twice)
//   - Graceful shutdown using context (optional stretch goal)
//
// The "web" is simulated by a map[string][]string where each key is a URL
// and the value is a list of URLs it links to. A fetch has simulated latency.
//
// Starting from "https://example.com/", discover all reachable pages.
//
// Expected behavior:
//   - All reachable URLs are printed exactly once.
//   - At most `maxConcurrency` fetches happen simultaneously.
//
// Hints:
//   - Use a buffered channel of capacity `maxConcurrency` as a semaphore:
//     send before fetch (acquire), receive after fetch (release).
//   - Track visited URLs with a map protected by sync.Mutex (or sync.Map).
//   - Use a sync.WaitGroup to know when all crawling is done.
//   - Recursive approach: for each discovered URL, if not visited, increment
//     the WaitGroup and launch a new goroutine.
//
// Stretch goal:
//   - Add context.WithTimeout to stop crawling after 2 seconds.
//
// Run:
//   go run .

package main

import (
	"fmt"
	"sync"
	"time"
)

// webGraph simulates the internet — each URL maps to the URLs it links to.
var webGraph = map[string][]string{
	"https://example.com/":         {"https://example.com/about", "https://example.com/blog"},
	"https://example.com/about":    {"https://example.com/team", "https://example.com/"},
	"https://example.com/blog":     {"https://example.com/blog/post1", "https://example.com/blog/post2"},
	"https://example.com/team":     {"https://example.com/about"},
	"https://example.com/blog/post1": {"https://example.com/blog"},
	"https://example.com/blog/post2": {"https://example.com/", "https://example.com/blog/post3"},
	"https://example.com/blog/post3": {},
}

// fetch simulates fetching a URL and returning its links.
func fetch(url string) []string {
	time.Sleep(100 * time.Millisecond) // simulate network latency
	return webGraph[url]
}

// TODO: Implement crawl — crawls the web graph starting from startURL.
// func crawl(startURL string, maxConcurrency int) { ... }

func main() {
	_ = fmt.Println
	_ = sync.Mutex{}
	_ = fetch

	// TODO: Crawl starting from "https://example.com/" with maxConcurrency = 3.
}
