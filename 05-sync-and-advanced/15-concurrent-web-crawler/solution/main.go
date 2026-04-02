// Solution 15: Concurrent Web Crawler
//
// Key takeaways:
//   - This problem composes almost every pattern from the previous exercises:
//       * Goroutines (Problem 1-2)
//       * WaitGroup for completion tracking (Problem 2)
//       * Mutex for shared state — the visited map (Problem 12)
//       * Buffered channel as a semaphore for bounded concurrency (Problem 10)
//       * Context for graceful shutdown (Problem 13)
//   - The semaphore pattern: a buffered channel of capacity N. Send to acquire
//     a slot (blocks if N goroutines are already working), receive to release.
//   - Deduplication must be checked-and-set atomically (under the lock), otherwise
//     two goroutines could both see a URL as unvisited and fetch it twice.
//   - The WaitGroup tracks in-flight work. We Add(1) before launching each
//     goroutine and Done() when it finishes, ensuring Wait() returns only when
//     the entire crawl tree has been explored.
//   - context.WithTimeout provides a hard deadline — crawling stops even if the
//     graph is enormous or has cycles that somehow bypass dedup.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var webGraph = map[string][]string{
	"https://example.com/":           {"https://example.com/about", "https://example.com/blog"},
	"https://example.com/about":      {"https://example.com/team", "https://example.com/"},
	"https://example.com/blog":       {"https://example.com/blog/post1", "https://example.com/blog/post2"},
	"https://example.com/team":       {"https://example.com/about"},
	"https://example.com/blog/post1": {"https://example.com/blog"},
	"https://example.com/blog/post2": {"https://example.com/", "https://example.com/blog/post3"},
	"https://example.com/blog/post3": {},
}

func fetch(url string) []string {
	time.Sleep(100 * time.Millisecond)
	return webGraph[url]
}

type crawler struct {
	visited map[string]bool
	mu      sync.Mutex
	sem     chan struct{} // semaphore for bounded concurrency
	wg      sync.WaitGroup
	ctx     context.Context
}

func newCrawler(maxConcurrency int, ctx context.Context) *crawler {
	return &crawler{
		visited: make(map[string]bool),
		sem:     make(chan struct{}, maxConcurrency),
		ctx:     ctx,
	}
}

func (c *crawler) visit(url string) {
	defer c.wg.Done()

	// Check for cancellation.
	select {
	case <-c.ctx.Done():
		return
	default:
	}

	// Acquire semaphore slot (bounded concurrency).
	c.sem <- struct{}{}
	links := fetch(url)
	<-c.sem // release

	fmt.Printf("crawled: %s (found %d links)\n", url, len(links))

	for _, link := range links {
		c.mu.Lock()
		alreadyVisited := c.visited[link]
		if !alreadyVisited {
			c.visited[link] = true
		}
		c.mu.Unlock()

		if !alreadyVisited {
			c.wg.Add(1)
			go c.visit(link)
		}
	}
}

func (c *crawler) crawl(startURL string) {
	c.mu.Lock()
	c.visited[startURL] = true
	c.mu.Unlock()

	c.wg.Add(1)
	go c.visit(startURL)
	c.wg.Wait()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c := newCrawler(3, ctx)
	c.crawl("https://example.com/")

	fmt.Printf("\nCrawled %d unique pages.\n", len(c.visited))
}
