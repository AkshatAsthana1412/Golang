// Problem 10: Worker Pool
//
// Implement a worker pool that processes jobs concurrently with a fixed number
// of workers.
//
// Setup:
//   - A Job is an int to be factorized (find all prime factors).
//   - A Result contains the original number and its factors.
//   - A dispatcher sends jobs onto a `jobs` channel.
//   - N worker goroutines read from `jobs`, compute factors, and send Results
//     to a `results` channel.
//   - Main collects and prints all results.
//
// Expected behavior:
//   - All numbers are factorized.
//   - Only N goroutines run at a time (bounded concurrency).
//
// Hints:
//   - Create two channels: jobs (chan int) and results (chan Result).
//   - Launch N workers, each running a loop: `for job := range jobs { ... }`.
//   - Send all jobs, then close the jobs channel so workers exit their loops.
//   - Collect len(jobs) results from the results channel.
//   - The `factorize` helper is provided.
//
// Run:
//   go run .

package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

// Result holds a number and its prime factors.
type Result struct {
	Number  int
	Factors []int
}

// factorize returns the prime factors of n.
func factorize(n int) []int {
	var factors []int
	d := 2
	for n > 1 {
		for n%d == 0 {
			factors = append(factors, d)
			n /= d
		}
		d++
	}
	return factors
}

// TODO: Implement worker — reads jobs from the jobs channel, factorizes, sends Result.
func worker(id int, jobs <-chan int, results chan<- Result) {
	for n := range jobs {
		results <- Result{Number: n, Factors: factorize(n)}
		// fmt.Printf("Worker %d, factorised number %d\n", id, n)
	}
}

func main() {
	_ = fmt.Println // remove once you use fmt

	// numbers := []int{12, 35, 77, 100, 131, 256, 360, 997, 1024, 2310}

	// TODO:
	// 1. Create jobs and results channels
	// 2. Launch N workers (try 3)
	// 3. Send all numbers as jobs, then close the jobs channel
	// 4. Collect and print all results
	start_time := time.Now()
	procs := runtime.GOMAXPROCS(0)
	fmt.Println("GOMAXPROCS:", procs)
	jobs := make(chan int)
	results := make(chan Result)
	var wg sync.WaitGroup

	// Get N from command line args, default to 3 if missing or invalid
	N := 3
	if len(os.Args) > 1 {
		var nArg int
		_, err := fmt.Sscanf(os.Args[1], "%d", &nArg)
		if err == nil && nArg > 0 {
			N = nArg
		} else {
			fmt.Println("Invalid N provided, using default N=3")
		}
	}
	fmt.Printf("Spawning %d workers\n", N)

	// spawning N workers
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, jobs, results)
		}(i)
	}

	// goroutine to generate jobs in the jobs channel
	go func() {
		defer close(jobs)
		for n := range 100000 {
			jobs <- n
		}
	}()

	// goroutine to close the results channel once all the workers are done processing
	go func() {
		wg.Wait()
		close(results)
	}()

	num_results := 0
	for range results {
		num_results++
	}

	fmt.Printf("Factorised %d numbers in %s.\n", num_results, time.Since(start_time))

	// Factorised 100000 numbers in 226.928584ms. 3 workers
	// Factorised 100000 numbers in 194.636083ms. 5 workers
	// Factorised 100000 numbers in 193.254125ms. 10 workers
	// Factorised 100000 numbers in 164.107959ms. 20 workers
}
