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

import "fmt"

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
// func worker(id int, jobs <-chan int, results chan<- Result) { ... }

func main() {
	_ = fmt.Println // remove once you use fmt

	numbers := []int{12, 35, 77, 100, 131, 256, 360, 997, 1024, 2310}

	// TODO:
	// 1. Create jobs and results channels
	// 2. Launch N workers (try 3)
	// 3. Send all numbers as jobs, then close the jobs channel
	// 4. Collect and print all results

	_ = numbers
}
