// Solution 14: Dining Philosophers
//
// Key takeaways:
//   - Deadlock requires four conditions (Coffman conditions): mutual exclusion,
//     hold and wait, no preemption, and circular wait. Breaking ANY one prevents
//     deadlock.
//   - Resource ordering breaks the "circular wait" condition: by always acquiring
//     the lower-numbered fork first, we guarantee that at least one philosopher
//     (the one between fork 0 and fork 4) acquires forks in reverse order from
//     their neighbor, breaking the cycle.
//   - sync.Mutex works well here because forks are naturally "owned resources"
//     rather than communication channels. This is a case where mutexes are more
//     idiomatic than channels.
//   - time.Sleep simulates eating/thinking and makes the output readable. In
//     real systems, this would be actual work.

package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numPhilosophers = 5
	numMeals        = 3
)

func philosopher(id int, forks [numPhilosophers]*sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	left := id
	right := (id + 1) % numPhilosophers

	// Resource ordering: always lock the lower-numbered fork first.
	first, second := left, right
	if right < left {
		first, second = right, left
	}

	for meal := 1; meal <= numMeals; meal++ {
		// Think
		fmt.Printf("philosopher %d is thinking\n", id)
		time.Sleep(50 * time.Millisecond)

		// Pick up forks (ordered)
		forks[first].Lock()
		forks[second].Lock()

		// Eat
		fmt.Printf("philosopher %d is eating (meal %d/%d)\n", id, meal, numMeals)
		time.Sleep(100 * time.Millisecond)

		// Put down forks
		forks[second].Unlock()
		forks[first].Unlock()
	}

	fmt.Printf("philosopher %d is done\n", id)
}

func main() {
	var forks [numPhilosophers]*sync.Mutex
	for i := range forks {
		forks[i] = &sync.Mutex{}
	}

	var wg sync.WaitGroup
	for i := range numPhilosophers {
		wg.Add(1)
		go philosopher(i, forks, &wg)
	}

	wg.Wait()
	fmt.Println("All philosophers have finished dining.")
}
