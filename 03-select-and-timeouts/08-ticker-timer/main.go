// Problem 8: Ticker and Timer
//
// Build a "heartbeat monitor" that uses time.NewTicker and time.NewTimer
// together in a select loop.
//
// Requirements:
//   - Create a Ticker that fires every 500ms to simulate periodic heartbeats.
//   - Create a Timer that fires after 2s to act as a hard deadline.
//   - In a select loop, handle both channels:
//       • On each tick, print "heartbeat at <elapsed time since start>"
//       • When the timer fires, print "deadline reached after <elapsed>" and exit
//   - Properly stop the ticker (defer ticker.Stop()) so it doesn't leak.
//
// Expected output (approximate):
//   heartbeat at 500ms
//   heartbeat at 1s
//   heartbeat at 1.5s
//   deadline reached after 2s
//
// Why NewTicker/NewTimer instead of time.Tick/time.After?
//   - time.Tick returns a bare channel with no way to stop the underlying
//     Ticker — it leaks resources. time.NewTicker returns a *Ticker whose
//     Stop() method releases the goroutine behind it.
//   - time.After is convenient in one-shot selects (like Problem 6), but
//     calling it inside a loop allocates a new Timer on every iteration
//     and none of them get garbage collected until they fire. NewTimer
//     lets you create one timer and optionally Reset it.
//
// Stretch goals:
//   1. After the loop, add a Reset to the timer (timer.Reset(1s)) and wait
//      for it again to see how timer reuse works.
//   2. Replace the fixed deadline with a "3 missed heartbeats" detector:
//      if no work arrives on a separate channel within 3 tick intervals,
//      print a warning using the ticker as the check cadence.
//
// Run:
//   go run .

package main

import (
	"fmt"
	"time"
)

func main() {
	// TODO:
	// 1. Record the start time with time.Now().
	// 2. Create a *time.Ticker with a 500ms interval.
	// 3. Defer ticker.Stop() to clean up the ticker goroutine.
	// 4. Create a *time.Timer with a 2s duration.
	// 5. Loop with select:
	//      case <-ticker.C:  print heartbeat with elapsed time
	//      case <-timer.C:   print deadline message and return

	_ = fmt.Println // remove once you use fmt
	_ = time.Now    // remove once you use time
}
