# Solutions & Concepts Guide

A detailed walkthrough of each problem's solution, the concurrency concepts it teaches, and common pitfalls to watch for.

---

## Level 1: Goroutine Basics

**Core concepts:** the `go` keyword, goroutine lifecycle, `sync.WaitGroup`

Goroutines are the foundation of Go concurrency. They are lightweight, user-space threads with ~2KB initial stacks, multiplexed onto OS threads by the Go runtime scheduler (M:N model). They are cheap to create — a program can comfortably run hundreds of thousands of them.

### Problem 1 — Hello Goroutine

**Directory:** `01-goroutines/01-hello-goroutine`

**What it teaches:** Launching goroutines with `go`, observing non-deterministic execution order, and understanding that `main` exiting kills all goroutines.

**Solution approach:**

A `for` loop launches 5 goroutines, each printing its index. A `time.Sleep` at the end of `main` gives them time to complete.

```go
for i := range 5 {
    go func() {
        fmt.Printf("Hello from goroutine %d\n", i)
    }()
}
time.Sleep(1 * time.Second)
```

**Key takeaways:**

- `go f()` starts `f` in a new goroutine. The calling goroutine does not wait for it to finish.
- Output order varies between runs because the scheduler interleaves goroutines non-deterministically.
- When `main()` returns, the process exits immediately — any goroutines still running are killed. The `time.Sleep` here is a deliberate hack; proper synchronization comes in Problem 2.

**Experiment:** Remove the `time.Sleep` and run several times. You'll see zero or partial output.

---

### Problem 2 — Wait For Goroutines

**Directory:** `01-goroutines/02-wait-for-goroutines`

**What it teaches:** Using `sync.WaitGroup` to wait for a batch of goroutines to finish.

**Solution approach:**

Each URL is fetched in its own goroutine. A `sync.WaitGroup` tracks completion: `Add(1)` before the launch, `defer wg.Done()` inside the goroutine, and `wg.Wait()` in `main`.

```go
var wg sync.WaitGroup

for _, url := range urls {
    wg.Add(1)
    go func() {
        defer wg.Done()
        result := simulateFetch(url)
        fmt.Println(result)
    }()
}

wg.Wait()
```

**Key takeaways:**

- `WaitGroup` is a counting semaphore: `Add(n)` increments, `Done()` decrements, `Wait()` blocks until zero.
- Always call `Add()` **before** launching the goroutine, not inside it. Otherwise there's a race where `Wait()` might see zero before the goroutine calls `Add()`.
- `defer wg.Done()` ensures the counter is decremented even if the goroutine panics.
- All fetches run concurrently, so total wall-clock time is roughly equal to the **slowest** fetch, not the sum.

**Pitfall:** Forgetting `wg.Add(1)` or mismatching the count causes either premature `Wait()` return or a permanent hang.

---

## Level 2: Channels

**Core concepts:** unbuffered channels, buffered channels, channel direction types (`chan<-`, `<-chan`), closing channels, `range` over channels

Channels are typed conduits for passing values between goroutines. An unbuffered channel creates a synchronization point (a "handoff") — both sender and receiver must be ready at the same instant. A buffered channel decouples them in time, up to the buffer's capacity.

### Problem 3 — Ping Pong

**Directory:** `02-channels/03-ping-pong`

**What it teaches:** Synchronous handoff on unbuffered channels to achieve strict turn-taking between goroutines.

**Solution approach:**

Two goroutines (ping and pong) share a single unbuffered `ball` channel. Each receives the ball, checks if the count exceeds 10, prints its name and the count, then sends back `count + 1`. A `done` channel signals `main` when the game ends.

```go
ball := make(chan int)
done := make(chan struct{})

go func() { // ping
    for {
        count := <-ball
        if count > 10 { done <- struct{}{}; return }
        fmt.Printf("ping: %d\n", count)
        ball <- count + 1
    }
}()

go func() { // pong
    // ... same structure ...
}()

ball <- 1 // serve
<-done
```

**Key takeaways:**

- An unbuffered channel enforces strict alternation: the sender blocks until the receiver is ready, creating a natural handshake.
- The `done` channel is idiomatic Go for signaling completion without `time.Sleep`.
- `struct{}` is a zero-size type — perfect for signal-only channels where the value itself doesn't matter.

---

### Problem 4 — Channel Directions

**Directory:** `02-channels/04-channel-directions`

**What it teaches:** Using directional channel types for compile-time safety and the pattern of each stage owning (and closing) its output channel.

**Solution approach:**

A `generator` function returns `<-chan int` (receive-only) and produces integers 1..N in a goroutine. A `squarer` function accepts `<-chan int`, reads values, squares them, and returns a new `<-chan int`. Both stages close their output channel when done.

```go
func generator(n int) <-chan int {
    out := make(chan int)
    go func() {
        for i := 1; i <= n; i++ { out <- i }
        close(out)
    }()
    return out
}

func squarer(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for v := range in { out <- v * v }
        close(out)
    }()
    return out
}
```

**Key takeaways:**

- `<-chan T` (receive-only) and `chan<- T` (send-only) enforce at **compile time** that a function only reads or writes. A bidirectional `chan T` is implicitly convertible to either.
- `close(ch)` signals "no more values." A `for v := range ch` loop exits when the channel is closed.
- **Convention:** the producer closes the channel, never the consumer. This avoids sending on a closed channel (which panics).

---

### Problem 5 — Buffered vs Unbuffered

**Directory:** `02-channels/05-buffered-vs-unbuffered`

**What it teaches:** The difference in blocking behavior between buffered and unbuffered channels, and how deadlocks arise.

**Solution approach:**

Three demonstrations:

1. **Buffered**: A `make(chan int, 3)` absorbs 3 sends without a receiver ready.
2. **Unbuffered deadlock**: A `make(chan int)` send blocks forever when no goroutine will receive — the runtime detects the deadlock and panics.
3. **Unbuffered fixed**: Launch a receiver goroutine before sending.

**Key takeaways:**

- Buffered channels decouple sender and receiver in time. Sends succeed immediately while the buffer has capacity.
- Unbuffered channels require both sides to rendezvous — a send blocks until a receive (and vice versa).
- Go's runtime detects when **all** goroutines are blocked with no possibility of progress and reports `fatal error: all goroutines are asleep - deadlock!`.
- Use buffered channels when you know a bounded amount of work will be produced before consumption. Use unbuffered when you want synchronization guarantees.

---

## Level 3: Select and Timeouts

**Core concepts:** `select` statement, `time.After`, `time.NewTicker`, `time.NewTimer`, multiplexing channel operations, the timeout pattern

`select` is Go's way of waiting on multiple channel operations simultaneously. It blocks until one case is ready, executes that case, and discards the rest. If multiple cases are ready, one is picked **at random**.

### Problem 6 — Timeout Fetch

**Directory:** `03-select-and-timeouts/06-timeout-fetch`

**What it teaches:** Using `select` with `time.After` to impose a deadline on a channel operation.

**Solution approach:**

A goroutine performs a simulated fetch (random 100-500ms latency) and sends the result on a channel. `main` uses `select` to race the result against `time.After(300ms)`.

```go
resultCh := make(chan string, 1)

go func() {
    resultCh <- simulateFetch()
}()

select {
case result := <-resultCh:
    fmt.Println("success:", result)
case <-time.After(300 * time.Millisecond):
    fmt.Println("timeout: fetch took too long")
}
```

**Key takeaways:**

- `time.After(d)` returns a `<-chan time.Time` that fires once after duration `d`. Combined with `select`, it creates a timeout.
- This pattern is fundamental to building resilient services — never wait forever for an operation that might hang.
- The timed-out goroutine continues running (the channel send will succeed into the buffer). In production, use `context.Context` for proper cancellation (see Problem 13).
- The channel is buffered with capacity 1 so the goroutine can complete its send even if `main` has already moved past the `select` (prevents goroutine leak).

---

### Problem 7 — First Response Wins

**Directory:** `03-select-and-timeouts/07-first-response-wins`

**What it teaches:** The hedged/redundant request pattern — send the same request to multiple replicas and accept whichever responds first.

**Solution approach:**

Three goroutines each simulate a replica with random latency and send their result on a shared channel. `main` receives once, getting whichever result arrives first.

```go
resultCh := make(chan string, len(replicas))

for _, name := range replicas {
    go func() {
        resultCh <- simulateReplica(name)
    }()
}

first := <-resultCh
```

**Key takeaways:**

- This pattern trades extra compute (redundant work) for reduced **tail latency** — you always get the fastest replica's response time.
- The channel **must** be buffered (capacity = number of goroutines). Without the buffer, the "losing" goroutines block forever on their send since nobody will receive their result. This is a **goroutine leak** — a common gotcha.
- This pattern is used at scale by companies like Google (see Jeff Dean's "tail at scale" paper).

---

### Problem 7b — Ticker and Timer

**Directory:** `03-select-and-timeouts/08-ticker-timer`

**What it teaches:** Using `time.NewTicker` and `time.NewTimer` for periodic and one-shot timed operations with proper resource cleanup — and why they should be preferred over `time.Tick` and `time.After` in loops.

**Solution approach:**

A `select` loop handles two time-based channels: a Ticker that fires every 500ms (heartbeats) and a Timer that fires once after 2s (deadline). The ticker is stopped via `defer` and the timer's channel naturally drains.

```go
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
```

**Key takeaways:**

- `time.NewTicker(d)` returns a `*Ticker` whose `.C` channel delivers ticks at regular intervals. Unlike `time.Tick(d)` (which returns a bare `<-chan time.Time`), the `*Ticker` has a `Stop()` method that releases the background goroutine. **Always stop tickers you create** — otherwise the goroutine behind them leaks for the lifetime of the process.
- `time.NewTimer(d)` returns a `*Timer` that fires once. Unlike `time.After(d)`, it exposes `Stop()` and `Reset(d)`. This matters in loops: calling `time.After` on every iteration creates a new timer each time, and unfired timers are **not garbage collected** until they expire — a classic memory leak in hot paths.
- Both `ticker.C` and `timer.C` are `<-chan time.Time`, so they compose naturally in a single `select` statement alongside any other channel operations. This "periodic work with a deadline" pattern is extremely common: health checks, metrics flushing, retry loops with timeouts.
- `timer.Reset(d)` can be used to reuse a timer for successive deadlines (e.g., idle timeout that resets on activity). **Caution:** `Reset` should only be called on stopped or expired timers. If the timer has already fired but the value hasn't been drained from `.C`, you must drain it first to avoid stale signals.

**When to use which:**

| Need | Use | Avoid |
|------|-----|-------|
| One-shot timeout in a non-loop `select` | `time.After(d)` | — |
| One-shot deadline in a loop or with reset | `time.NewTimer(d)` | `time.After` (leaks per iteration) |
| Periodic operation, short-lived | `time.NewTicker(d)` + `Stop()` | `time.Tick` (leaks) |
| Periodic operation, process lifetime | `time.Tick(d)` is acceptable | — |

**Pitfall:** Problem 11 (Rate Limiter) deliberately uses `time.Tick` and notes the caveat. This problem shows the proper alternative — if you ever need to stop rate-limiting, you need `NewTicker` so you can call `Stop()`.

---

## Level 4: Classic Concurrency Patterns

**Core concepts:** pipeline, fan-out/fan-in, worker pool, rate limiting with channels

These are the bread-and-butter concurrency patterns in Go. They compose the primitives from earlier levels into reusable architectures.

### Problem 8 — Pipeline

**Directory:** `04-patterns/08-pipeline`

**What it teaches:** Building a multi-stage processing pipeline where each stage is a goroutine connected by channels.

**Solution approach:**

Three stages: `generate(nums)` emits integers, `square(in)` squares them, `filterOdd(in)` keeps only odd values. Each stage reads from an input channel and writes to an output channel, closing it when done. They compose naturally:

```go
ch := filterOdd(square(generate(nums)))

for v := range ch {
    fmt.Println(v)
}
```

**Key takeaways:**

- A **pipeline** is a series of stages connected by channels. Each stage owns and closes its output channel.
- **Backpressure is automatic:** if a downstream stage is slow, the upstream stage blocks on its channel send. No explicit flow-control code is needed.
- Pipelines compose like function calls — you can add, remove, or reorder stages by changing how channels are wired.
- Shutdown propagates naturally: when a stage closes its output channel, the downstream `range` loop exits.

---

### Problem 9 — Fan-Out / Fan-In

**Directory:** `04-patterns/09-fan-out-fan-in`

**What it teaches:** Parallelizing a pipeline stage across multiple workers (fan-out) and merging their outputs (fan-in).

**Solution approach:**

The `generate` stage feeds a single channel. Three `square` workers all read from the same channel (fan-out — the runtime delivers each value to exactly one receiver). A `merge` function combines the three output channels into one using a forwarding goroutine per channel and a `WaitGroup` to close the merged output.

```go
func merge(channels ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup

    for _, ch := range channels {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for v := range ch { out <- v }
        }()
    }

    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}
```

**Key takeaways:**

- **Fan-out:** multiple goroutines read from one channel. Go guarantees each value is delivered to exactly one receiver — no duplicates, no extra synchronization needed.
- **Fan-in:** a merge goroutine forwards values from N input channels to one output channel. A `WaitGroup` ensures the output channel is closed only after all inputs are drained.
- Output order is **non-deterministic** because workers race to process items.
- This pattern is ideal for CPU-bound pipeline stages where you want to utilize multiple cores.

---

### Problem 10 — Worker Pool

**Directory:** `04-patterns/10-worker-pool`

**What it teaches:** Bounded concurrency through a fixed number of workers reading from a shared job queue.

**Solution approach:**

A `jobs` channel carries integers to be factorized. N worker goroutines read from `jobs`, compute factors, and send results to a `results` channel. `main` sends all jobs, closes the channel, and collects all results.

```go
func worker(id int, jobs <-chan int, results chan<- Result) {
    for n := range jobs {
        factors := factorize(n)
        results <- Result{Number: n, Factors: factors}
    }
}
```

**Key takeaways:**

- The worker pool pattern **bounds concurrency**: no matter how many jobs exist, only N workers run simultaneously. This prevents resource exhaustion.
- Workers share a single `jobs` channel. The runtime naturally **load-balances** — faster workers pick up more jobs.
- Closing the `jobs` channel causes all workers' `range` loops to exit cleanly — this is the shutdown signal.
- The `results` channel is buffered to `len(numbers)` to avoid deadlock: if workers filled an unbuffered channel faster than `main` drained it, the system would stall.

**Pitfall:** Sending jobs on the same goroutine that also reads results from an unbuffered channel can deadlock. Either buffer the channels or send jobs in a separate goroutine.

---

### Problem 11 — Rate Limiter

**Directory:** `04-patterns/11-rate-limiter`

**What it teaches:** Channel-based rate limiting using `time.Tick` for steady throughput and a buffered channel as a token bucket for burst support.

**Solution approach:**

**Part A — Steady rate:** `time.Tick(200ms)` returns a channel that fires every 200ms. Receiving from it before processing each request limits throughput to 5/second.

**Part B — Burst + steady rate:** A buffered channel of capacity 3 is pre-filled with tokens. A background goroutine refills it at 200ms intervals. Consumers drain tokens immediately during a burst, then fall back to the steady refill rate.

```go
burstyBucket := make(chan struct{}, 3)
for range 3 { burstyBucket <- struct{}{} } // pre-fill

go func() {
    for range time.Tick(200 * time.Millisecond) {
        select {
        case burstyBucket <- struct{}{}:
        default: // bucket full, discard
        }
    }
}()

for _, req := range requests {
    <-burstyBucket // blocks when no tokens available
    process(req)
}
```

**Key takeaways:**

- `time.Tick(d)` delivers a value every `d` — blocking on it before each operation naturally throttles throughput.
- A **token bucket** is a buffered channel pre-loaded with tokens. Burst capacity = buffer size. The `select` with `default` in the refiller prevents blocking when the bucket is already full.
- This is the same fundamental algorithm behind production rate limiters like `golang.org/x/time/rate`.
- **Caveat:** `time.Tick` leaks the underlying ticker. In long-running programs, use `time.NewTicker` and call `ticker.Stop()`.

---

## Level 5: Sync Primitives and Advanced Problems

**Core concepts:** `sync.Mutex`, data races, `context.Context` cancellation, cooperative shutdown, composing all prior patterns

This level shifts from pure channel-based communication to scenarios that require shared-state protection and coordinated cancellation across goroutine trees.

### Problem 12 — Safe Counter

**Directory:** `05-sync-and-advanced/12-safe-counter`

**What it teaches:** Understanding data races, fixing them with `sync.Mutex`, and comparing the mutex approach to a channel-based approach.

**Solution approach:**

Three implementations of concurrent counter increment:

1. **Racy** — multiple goroutines do `counter++` with no synchronization. The read-modify-write is not atomic, so the final count is wrong and `go run -race` reports a data race.
2. **Mutex** — wraps each `counter++` in `mu.Lock()` / `mu.Unlock()`. Correct and simple.
3. **Channel** — goroutines send `1` on a channel; a single goroutine owns the counter and sums all received values. No shared state, no race.

```go
// Mutex approach
mu.Lock()
counter++
mu.Unlock()

// Channel approach
for v := range ch {
    total += v  // single owner, no race
}
```

**Key takeaways:**

- A **data race** occurs when two goroutines access the same variable concurrently and at least one is a write. Go's race detector (`go run -race`) catches this at runtime.
- `sync.Mutex` provides mutual exclusion: only one goroutine holds the lock at a time.
- The channel approach eliminates shared state entirely — it embodies the Go proverb: *"share memory by communicating."*
- **When to use which:** Mutex is simpler for "protect a variable" scenarios. Channels are better when the interaction is naturally about communication or pipelines.

---

### Problem 13 — Context Cancellation

**Directory:** `05-sync-and-advanced/13-context-cancellation`

**What it teaches:** Cooperative cancellation of goroutines using `context.WithCancel`, and the `ctx.Done()` pattern.

**Solution approach:**

Three goroutines search different "databases" with random latency. Each checks `ctx.Done()` inside a `select` to detect cancellation. When the first result arrives, `main` calls `cancel()` to signal the others to stop.

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// In each search goroutine:
select {
case <-ctx.Done():
    return "", ctx.Err()  // cancelled
case <-time.After(latency):
    return result, nil    // completed
}
```

**Key takeaways:**

- `context.WithCancel` returns a derived context and a `cancel` function. Calling `cancel()` closes `ctx.Done()`, which all watching goroutines can detect.
- Cancellation is **cooperative**: goroutines must explicitly check `ctx.Done()` (usually via `select`). If they don't check, `cancel()` has no effect.
- Always `defer cancel()` to prevent context leaks, even if you call it explicitly earlier.
- The pattern generalizes: `context.WithTimeout` and `context.WithDeadline` cancel automatically after a duration or at a specific time. These are used pervasively in production Go — HTTP handlers, gRPC calls, database queries all accept a `context.Context`.

---

### Problem 14 — Dining Philosophers *(not yet implemented)*

**Directory:** `05-sync-and-advanced/14-dining-philosophers`

**What it will teach:** Deadlock avoidance through asymmetric resource ordering.

**Planned concepts:**

- 5 philosophers, 5 forks modeled as channels or mutexes
- **Deadlock** occurs when each philosopher holds one fork and waits for the other — circular wait
- **Resource ordering** breaks the cycle: one philosopher picks up forks in reverse order, preventing all from waiting simultaneously
- Understanding of the four Coffman conditions for deadlock: mutual exclusion, hold and wait, no preemption, circular wait

---

### Problem 15 — Concurrent Web Crawler *(not yet implemented)*

**Directory:** `05-sync-and-advanced/15-concurrent-web-crawler`

**What it will teach:** Composing all prior patterns into a realistic program.

**Planned concepts:**

- Bounded parallelism via a semaphore (buffered channel)
- Deduplication using `sync.Map` or a mutex-guarded map
- Graceful shutdown via `context.Context`
- Recursive task spawning with proper synchronization
- Combines: goroutines, channels, WaitGroup, mutex, context, worker pool, fan-out

---

## Concepts Summary by Level

| Level | Problems | New Concepts | Go Primitives Used |
|-------|----------|-------------|--------------------|
| 1 | 1–2 | Goroutine lifecycle, waiting for completion | `go`, `sync.WaitGroup` |
| 2 | 3–5 | Channel communication, blocking, direction types | `chan`, `<-chan`, `chan<-`, `close`, `range` |
| 3 | 6–7b | Multiplexing channels, timeouts, redundant requests, periodic ticks | `select`, `time.After`, `time.NewTicker`, `time.NewTimer` |
| 4 | 8–11 | Pipeline, fan-out/fan-in, worker pool, rate limiting | Channel composition, `time.Tick`, token bucket |
| 5 | 12–15 | Shared state, races, cancellation, deadlock avoidance | `sync.Mutex`, `context.Context`, race detector |

---

## Common Pitfalls Reference

| Pitfall | Where It Appears | Fix |
|---------|-----------------|-----|
| Main exits before goroutines finish | Problem 1 | Use `sync.WaitGroup` (Problem 2) |
| Calling `wg.Add()` inside the goroutine | Problem 2 | Always call `Add()` before `go` |
| Sending on a closed channel (panics) | Problem 4 | Only the **producer** should close |
| Goroutine leak from unbuffered channel send | Problem 7 | Buffer the channel to match sender count |
| Data race on shared variable | Problem 12 | Use `sync.Mutex` or channel ownership |
| Deadlock from unbuffered channel in single goroutine | Problem 5 | Use a separate goroutine or buffer |
| Forgetting to check `ctx.Done()` | Problem 13 | Always use `select` with `ctx.Done()` case |
| `time.Tick` ticker leak | Problems 7b, 11 | Use `time.NewTicker` + `defer ticker.Stop()` |
| `time.After` in a loop (memory leak) | Problem 7b | Use `time.NewTimer` + `Reset()` |

---

## Running Solutions with the Race Detector

Every solution should be run with Go's race detector enabled:

```bash
cd <problem-directory>/solution
go run -race .
```

The race detector instruments memory accesses at compile time and reports concurrent unsynchronized access at runtime. It has no false positives — if it reports a race, there is one. Problem 12 is specifically designed to trigger it.

---

## Design Philosophy

The solutions follow a consistent philosophy drawn from Go's concurrency ethos:

1. **Prefer channels for communication, mutexes for state protection.** When goroutines need to pass data or signal events, channels are clearer. When they just need to protect a shared variable, a mutex is simpler.

2. **The producer closes the channel.** This is a convention throughout all solutions. The entity that writes to a channel is responsible for closing it. Consumers just `range` over it.

3. **Compose small, focused stages.** Each pipeline stage, worker, or utility function does one thing. Complex behavior emerges from wiring them together with channels.

4. **Always handle shutdown.** Every solution demonstrates clean termination — via `WaitGroup`, channel closing, or `context` cancellation. No goroutine is left running by accident.

5. **Use the race detector.** It's not optional. Run `go run -race .` on every concurrent program during development.
