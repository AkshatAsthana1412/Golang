// Problem 3: Ping Pong
//
// Two goroutines ("ping" and "pong") pass a ball back and forth.
// The ball is an integer counter that increments with each pass.
// The game ends when the counter reaches 10.
//
// Expected output (exact):
//   ping: 1
//   pong: 2
//   ping: 3
//   pong: 4
//   ...
//   pong: 10
//
// Hints:
//   - Use an unbuffered channel for each direction (ping→pong and pong→ping),
//     or a single channel if you prefer.
//   - An unbuffered channel send blocks until the other side receives — this is
//     what enforces the strict alternation.
//   - One goroutine must initiate by sending the first value.
//
// Run:
//   go run .

package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	// TODO: Create channel(s) and launch two goroutines (ping and pong).
	// Each receives the ball, increments it, prints its role and the count,
	// and sends the ball to the other side.
	// Stop when the count reaches 10.
	// Trace must go to its own file (or stderr); mixing trace binary data with
	// fmt.Println on stdout corrupts the file (go tool trace will fail to parse).
	traceFile, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer traceFile.Close()
	if err := trace.Start(traceFile); err != nil {
		panic(err)
	}
	defer trace.Stop()
	ball := make(chan int)
	done := make(chan struct{})

	// ping
	go func() {
		for {
			count := <-ball
			if count > 10 {
				done <- struct{}{}
				return
			}
			ball <- count + 1
			fmt.Println("Ball hit by ping.")
		}
	}()

	// pong
	go func() {
		for {
			count := <-ball
			if count > 10 {
				done <- struct{}{}
				return
			}
			ball <- count + 1
			fmt.Println("Ball hit by pong.")
		}
	}()

	ball <- 1
	<-done
}
