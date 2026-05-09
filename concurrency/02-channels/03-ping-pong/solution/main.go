// Solution 3: Ping Pong
//
// Key takeaways:
//   - An unbuffered channel is a synchronization point: the sender blocks until
//     the receiver is ready, and vice versa. This gives us strict alternation
//     without any explicit locks.
//   - We use a single channel here. The "ping" goroutine sends first, then
//     the "pong" goroutine receives, increments, and sends back.
//   - The `done` channel signals main that the game is over so it can exit
//     cleanly instead of using time.Sleep.

package main

import "fmt"

func main() {
	ball := make(chan int)
	done := make(chan struct{})

	// Ping player
	go func() {
		for {
			count := <-ball
			if count > 10 {
				done <- struct{}{}
				return
			}
			fmt.Printf("ping: %d\n", count)
			ball <- count + 1
		}
	}()

	// Pong player
	go func() {
		for {
			count := <-ball
			if count > 10 {
				done <- struct{}{}
				return
			}
			fmt.Printf("pong: %d\n", count)
			ball <- count + 1
		}
	}()

	// Serve: send the first ball to kick off the game.
	ball <- 1

	<-done
}
