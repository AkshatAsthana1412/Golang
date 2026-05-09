// Solution 10: Worker Pool
//
// Key takeaways:
//   - The worker pool pattern bounds concurrency: no matter how many jobs you
//     have, only N workers process them simultaneously.
//   - Workers share a single `jobs` channel. When a job is available, exactly
//     one worker picks it up (channel receive semantics). This is automatic
//     load balancing — faster workers naturally get more jobs.
//   - Closing the `jobs` channel causes all workers' `range` loops to exit,
//     which is the clean shutdown signal.
//   - Sending jobs in a goroutine avoids deadlock: if jobs is unbuffered (or
//     has limited buffer), the main goroutine would block on send once the
//     buffer is full and workers haven't started consuming yet.

package main

import "fmt"

type Result struct {
	Number  int
	Factors []int
}

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

func worker(id int, jobs <-chan int, results chan<- Result) {
	for n := range jobs {
		factors := factorize(n)
		fmt.Printf("worker %d: factorized %d\n", id, n)
		results <- Result{Number: n, Factors: factors}
	}
}

func main() {
	numbers := []int{12, 35, 77, 100, 131, 256, 360, 997, 1024, 2310}
	numWorkers := 3

	jobs := make(chan int, len(numbers))
	results := make(chan Result, len(numbers))

	// Launch workers.
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs.
	for _, n := range numbers {
		jobs <- n
	}
	close(jobs)

	// Collect results.
	for range len(numbers) {
		r := <-results
		fmt.Printf("%d = %v\n", r.Number, r.Factors)
	}
}
