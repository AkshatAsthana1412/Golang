# Go Concurrency Patterns

A hands-on collection of 16 progressively challenging problems designed to teach Go concurrency from first principles.

## How to Use

Each problem lives in its own directory with two files:

- **`main.go`** — A starter scaffold with the problem statement, hints, and function signatures. **This is the file you edit.**
- **`solution/main.go`** — A reference solution with detailed comments. Read this *after* you've attempted the problem yourself.

Work through the problems **in order** (1 through 15). Each one builds on concepts from the previous.

### Running a Problem

```bash
cd 01-goroutines/01-hello-goroutine
go run .            # run your solution
go run -race .      # run with the race detector enabled
```

### Workflow for Each Problem

1. Read the problem statement and hints in `main.go`
2. Implement the solution
3. Run with `go run -race .` to check for data races
4. Compare with the reference in `solution/main.go`
5. Experiment — change buffer sizes, worker counts, timeouts — observe behavior changes

---

## Problem Index

| # | Problem | Directory | Key Concepts |
|---|---------|-----------|--------------|
| 1 | Hello Goroutine | `01-goroutines/01-hello-goroutine` | `go` keyword, goroutine lifecycle |
| 2 | Wait For Goroutines | `01-goroutines/02-wait-for-goroutines` | `sync.WaitGroup` |
| 3 | Ping Pong | `02-channels/03-ping-pong` | unbuffered channels, synchronous handoff |
| 4 | Channel Directions | `02-channels/04-channel-directions` | `chan<-`, `<-chan`, closing channels |
| 5 | Buffered vs Unbuffered | `02-channels/05-buffered-vs-unbuffered` | buffered channels, blocking semantics |
| 6 | Timeout Fetch | `03-select-and-timeouts/06-timeout-fetch` | `select`, `time.After` |
| 7 | First Response Wins | `03-select-and-timeouts/07-first-response-wins` | redundant requests, `select` |
| 7b | Ticker and Timer | `03-select-and-timeouts/08-ticker-timer` | `time.NewTicker`, `time.NewTimer`, resource cleanup |
| 8 | Pipeline | `04-patterns/08-pipeline` | channel chaining, backpressure |
| 9 | Fan-Out / Fan-In | `04-patterns/09-fan-out-fan-in` | parallel workers, merge |
| 10 | Worker Pool | `04-patterns/10-worker-pool` | bounded concurrency, job queues |
| 11 | Rate Limiter | `04-patterns/11-rate-limiter` | `time.Tick`, token bucket |
| 12 | Safe Counter | `05-sync-and-advanced/12-safe-counter` | `sync.Mutex`, race detector |
| 13 | Context Cancellation | `05-sync-and-advanced/13-context-cancellation` | `context.WithCancel`, `ctx.Done()` |
| 14 | Dining Philosophers | `05-sync-and-advanced/14-dining-philosophers` | deadlock avoidance, resource ordering |
| 15 | Concurrent Web Crawler | `05-sync-and-advanced/15-concurrent-web-crawler` | composition of all patterns |

---

## Quick Concepts Reference

### Goroutines

Lightweight, user-space threads managed by the Go runtime. Each starts with ~2KB of stack (grows as needed) and is multiplexed onto a small number of OS threads by the Go scheduler (M:N threading). Launching one is as simple as `go f()`.

### Channels

Typed conduits for communication between goroutines.

- **Unbuffered** (`make(chan T)`) — sender blocks until a receiver is ready, and vice versa. This creates a synchronization point (a "handoff").
- **Buffered** (`make(chan T, n)`) — sender blocks only when the buffer is full; receiver blocks only when the buffer is empty.
- Close a channel with `close(ch)` to signal that no more values will be sent. Receivers can detect this with the `v, ok := <-ch` idiom or by `range`-ing over the channel.

### Select

Multiplexes across multiple channel operations. If multiple cases are ready simultaneously, one is chosen **at random** (not in order). A `default` case makes the select non-blocking.

```go
select {
case v := <-ch1:
    // ch1 is ready
case ch2 <- val:
    // sent to ch2
case <-time.After(1 * time.Second):
    // timeout
default:
    // non-blocking fallthrough
}
```

### sync.WaitGroup

A counting barrier. Call `Add(n)` before launching goroutines, `Done()` (which decrements by 1) at the end of each goroutine, and `Wait()` to block until the count reaches zero.

### sync.Mutex / RWMutex

Classic mutual exclusion locks for protecting shared state. `RWMutex` allows multiple concurrent readers but exclusive writers. Prefer channels when the interaction is naturally "communication"; prefer mutexes when it's naturally "shared state protection."

### context.Context

Carries deadlines, cancellation signals, and request-scoped values across goroutine trees. Key constructors: `context.WithCancel`, `context.WithTimeout`, `context.WithDeadline`. Goroutines cooperatively check `ctx.Done()` to stop work early.

### The Go Proverb

> Don't communicate by sharing memory; share memory by communicating.

Use channels to pass data ownership between goroutines rather than protecting shared data with locks, when the design allows it.
