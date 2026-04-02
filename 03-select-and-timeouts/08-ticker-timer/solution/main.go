// Solution 8: Ticker and Timer
//
// Key takeaways:
//   - time.NewTicker returns a *Ticker that fires repeatedly at a fixed interval.
//     Always call ticker.Stop() when done — otherwise the backing goroutine leaks.
//     The convenience function time.Tick omits Stop and should only be used in
//     programs that run for the entire process lifetime (e.g., a main loop that
//     never returns).
//   - time.NewTimer returns a *Timer that fires once after the given duration.
//     Unlike time.After (which you can't stop or reset), NewTimer exposes Stop()
//     and Reset(d) for reuse. This matters in loops: calling time.After on every
//     iteration creates a new timer each time, and timers that haven't fired yet
//     are not garbage collected — a classic memory leak in hot paths.
//   - ticker.C and timer.C are both <-chan time.Time, so they compose naturally
//     in a single select statement alongside any other channel operations.
//   - The select loop here is a common real-world pattern: do periodic work
//     (health checks, metrics flush, retry loops) until a deadline or
//     cancellation signal arrives.

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	deadline := time.NewTimer(2 * time.Second)

	for {
		select {
		case t := <-ticker.C:
			fmt.Printf("heartbeat at %v\n", t.Sub(start).Round(time.Millisecond))
		case t := <-deadline.C:
			fmt.Printf("deadline reached after %v\n", t.Sub(start).Round(time.Millisecond))
			return
		}
	}
}
