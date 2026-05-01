# Go Context Guide for Concurrent Applications

## Table of Contents
- [Context Types](#context-types)
- [Context Hierarchies](#context-hierarchies)  
- [Value Propagation](#value-propagation)
- [Common Patterns](#common-patterns)
- [Testing with Context](#testing-with-context)
- [Best Practices](#best-practices)
- [Anti-Patterns to Avoid](#anti-patterns-to-avoid)
- [Real-World Examples](#real-world-examples)

## Context Types

### 1. Background Context
```go
ctx := context.Background()
```
- **Use**: Root context for main functions, initialization, tests
- **Properties**: Never cancelled, no deadline, no values
- **When**: Starting point for context chains

### 2. TODO Context  
```go
ctx := context.TODO()
```
- **Use**: Placeholder when you're unsure what context to use
- **Properties**: Same as Background but signals "needs proper context later"
- **When**: Refactoring code, temporary implementations

### 3. Cancellation Context
```go
ctx, cancel := context.WithCancel(parent)
defer cancel() // Always cleanup!
```
- **Use**: Manual cancellation control
- **Triggers**: Explicit `cancel()` call
- **Example**: Cancel API calls when user navigates away

### 4. Timeout Context
```go
ctx, cancel := context.WithTimeout(parent, 5*time.Second)
defer cancel()
```
- **Use**: Automatic cancellation after duration
- **Triggers**: After specified timeout OR explicit `cancel()`
- **Example**: HTTP client timeouts, database query limits

### 5. Deadline Context
```go
deadline := time.Now().Add(10*time.Minute)
ctx, cancel := context.WithDeadline(parent, deadline)
defer cancel()
```
- **Use**: Cancellation at specific time
- **Triggers**: When deadline reached OR explicit `cancel()`
- **Example**: Request must complete before business hours end

### 6. Value Context
```go
ctx := context.WithValue(parent, "userID", 12345)
ctx = context.WithValue(ctx, "requestID", "abc-123")
```
- **Use**: Pass request-scoped data
- **Caution**: Use sparingly, prefer explicit parameters
- **Example**: User authentication, request tracing

## Context Hierarchies

### Parent-Child Relationships
```go
// Root context
root := context.Background()

// Service-level timeout (30s max)
serviceCtx, serviceCancel := context.WithTimeout(root, 30*time.Second)
defer serviceCancel()

// Request-level timeout (5s max, but inherits 30s limit)
requestCtx, requestCancel := context.WithTimeout(serviceCtx, 5*time.Second)
defer requestCancel()

// Add request ID
requestCtx = context.WithValue(requestCtx, "requestID", generateID())
```

### Cancellation Propagation
- **Downstream**: Cancelling parent cancels ALL children
- **Upstream**: Cancelling child does NOT cancel parent
- **Immediate**: Cancellation propagates instantly to all descendants

```go
parent, parentCancel := context.WithCancel(context.Background())
child, childCancel := context.WithCancel(parent)

parentCancel() // child is also cancelled
// childCancel() // would only cancel child
```

## Value Propagation

### Type-Safe Keys
```go
type contextKey string

const (
    UserIDKey    contextKey = "userID"
    RequestIDKey contextKey = "requestID"
)

// Setting values
ctx := context.WithValue(parent, UserIDKey, 12345)

// Getting values
userID, ok := ctx.Value(UserIDKey).(int)
if !ok {
    // Handle missing or wrong type
}
```

### Value Helper Functions
```go
func WithUserID(ctx context.Context, userID int) context.Context {
    return context.WithValue(ctx, UserIDKey, userID)
}

func GetUserID(ctx context.Context) (int, bool) {
    userID, ok := ctx.Value(UserIDKey).(int)
    return userID, ok
}

func MustGetUserID(ctx context.Context) int {
    userID, ok := GetUserID(ctx)
    if !ok {
        panic("userID not found in context")
    }
    return userID
}
```

## Common Patterns

### 1. Fan-Out with Cancellation
```go
func searchMultipleSources(ctx context.Context, query string) (string, error) {
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    results := make(chan string, 3)
    errors := make(chan error, 3)

    sources := []string{"db1", "db2", "cache"}
    
    for _, source := range sources {
        go func(src string) {
            result, err := searchSource(ctx, src, query)
            if err != nil {
                errors <- err
                return
            }
            select {
            case results <- result:
                cancel() // Cancel others on first success
            case <-ctx.Done():
                return
            }
        }(source)
    }

    select {
    case result := <-results:
        return result, nil
    case <-ctx.Done():
        return "", ctx.Err()
    }
}
```

### 2. Pipeline with Context
```go
func pipeline(ctx context.Context, input <-chan int) <-chan string {
    output := make(chan string)
    
    go func() {
        defer close(output)
        for {
            select {
            case <-ctx.Done():
                return
            case num, ok := <-input:
                if !ok {
                    return
                }
                // Process with context check
                processed := processWithContext(ctx, num)
                select {
                case output <- processed:
                case <-ctx.Done():
                    return
                }
            }
        }
    }()
    
    return output
}
```

### 3. Retry with Backoff
```go
func retryWithContext(ctx context.Context, fn func() error) error {
    backoff := time.Millisecond * 100
    maxBackoff := time.Second * 10
    
    for attempt := 0; attempt < 5; attempt++ {
        if err := fn(); err == nil {
            return nil
        }
        
        select {
        case <-ctx.Done():
            return ctx.Err()
        case <-time.After(backoff):
            backoff *= 2
            if backoff > maxBackoff {
                backoff = maxBackoff
            }
        }
    }
    return fmt.Errorf("max retries exceeded")
}
```

### 4. Worker Pool
```go
func workerPool(ctx context.Context, jobs <-chan Job, numWorkers int) <-chan Result {
    results := make(chan Result)
    
    var wg sync.WaitGroup
    
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            for {
                select {
                case <-ctx.Done():
                    return
                case job, ok := <-jobs:
                    if !ok {
                        return
                    }
                    result := processJob(ctx, job)
                    select {
                    case results <- result:
                    case <-ctx.Done():
                        return
                    }
                }
            }
        }(i)
    }
    
    go func() {
        wg.Wait()
        close(results)
    }()
    
    return results
}
```

## Testing with Context

### 1. Testing Cancellation
```go
func TestFunctionCancellation(t *testing.T) {
    ctx, cancel := context.WithCancel(context.Background())
    
    done := make(chan bool)
    go func() {
        err := longRunningFunction(ctx)
        assert.Equal(t, context.Canceled, err)
        done <- true
    }()
    
    time.Sleep(100 * time.Millisecond)
    cancel()
    
    select {
    case <-done:
        // Success
    case <-time.After(time.Second):
        t.Fatal("function didn't respond to cancellation")
    }
}
```

### 2. Testing Timeouts
```go
func TestTimeout(t *testing.T) {
    ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
    defer cancel()
    
    start := time.Now()
    err := slowFunction(ctx)
    elapsed := time.Since(start)
    
    assert.Equal(t, context.DeadlineExceeded, err)
    assert.True(t, elapsed < 100*time.Millisecond, "should timeout quickly")
}
```

### 3. Mock Context Values
```go
func TestWithMockUser(t *testing.T) {
    ctx := context.WithValue(context.Background(), UserIDKey, 12345)
    
    result, err := userSpecificFunction(ctx)
    
    assert.NoError(t, err)
    assert.Equal(t, "user-12345-data", result)
}
```

## Best Practices

### 1. Function Signatures
```go
// ✅ Good - context first parameter
func ProcessData(ctx context.Context, data []byte) error

// ❌ Bad - context not first
func ProcessData(data []byte, ctx context.Context) error

// ✅ Good - return context from constructors
func NewService(cfg Config) (context.Context, *Service, error)
```

### 2. Context Checking Frequency
```go
// ✅ Good - check context in loops
func processLargeDataset(ctx context.Context, data []Record) error {
    for i, record := range data {
        if i%100 == 0 { // Check every 100 iterations
            select {
            case <-ctx.Done():
                return ctx.Err()
            default:
            }
        }
        
        if err := processRecord(record); err != nil {
            return err
        }
    }
    return nil
}
```

### 3. Resource Cleanup
```go
func connectToDatabase(ctx context.Context) (*sql.DB, error) {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel() // Always cleanup
    
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }
    
    // Test connection with context
    if err := db.PingContext(ctx); err != nil {
        db.Close()
        return nil, err
    }
    
    return db, nil
}
```

## Anti-Patterns to Avoid

### 1. Don't Store Context in Structs
```go
// ❌ Bad
type Service struct {
    ctx context.Context // Don't do this!
    db  *sql.DB
}

// ✅ Good - pass context through methods
type Service struct {
    db *sql.DB
}

func (s *Service) ProcessData(ctx context.Context, data []byte) error {
    return s.db.ExecContext(ctx, "INSERT ...", data)
}
```

### 2. Don't Pass nil Context
```go
// ❌ Bad
ProcessData(nil, data)

// ✅ Good
ProcessData(context.Background(), data)
ProcessData(context.TODO(), data) // If unsure
```

### 3. Don't Overuse Context Values
```go
// ❌ Bad - too many values
ctx = context.WithValue(ctx, "param1", val1)
ctx = context.WithValue(ctx, "param2", val2)
ctx = context.WithValue(ctx, "param3", val3)

// ✅ Good - use parameters
func ProcessData(ctx context.Context, param1, param2, param3 string) error
```

### 4. Don't Ignore Context Errors
```go
// ❌ Bad - ignoring cancellation
func badFunction(ctx context.Context) error {
    for i := 0; i < 1000; i++ {
        doWork(i)
    }
    return nil
}

// ✅ Good - respecting cancellation
func goodFunction(ctx context.Context) error {
    for i := 0; i < 1000; i++ {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            doWork(i)
        }
    }
    return nil
}
```

## Real-World Examples

### 1. HTTP Server with Context
```go
func httpHandler(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
    defer cancel()
    
    // Add request ID for tracing
    requestID := generateRequestID()
    ctx = context.WithValue(ctx, RequestIDKey, requestID)
    
    // Pass to business logic
    result, err := businessLogic(ctx, r.Body)
    if err != nil {
        if errors.Is(err, context.DeadlineExceeded) {
            http.Error(w, "Request timeout", http.StatusRequestTimeout)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    json.NewEncoder(w).Encode(result)
}
```

### 2. Database with Context
```go
func getUserByID(ctx context.Context, db *sql.DB, userID int) (*User, error) {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    
    query := "SELECT id, name, email FROM users WHERE id = $1"
    row := db.QueryRowContext(ctx, query, userID)
    
    var user User
    err := row.Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        if errors.Is(err, context.DeadlineExceeded) {
            return nil, fmt.Errorf("database query timeout: %w", err)
        }
        return nil, err
    }
    
    return &user, nil
}
```

### 3. Microservice Communication
```go
func callExternalService(ctx context.Context, url string, payload interface{}) (*Response, error) {
    // Inherit parent timeout but add buffer
    ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()
    
    // Add tracing headers
    req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
    if err != nil {
        return nil, err
    }
    
    // Propagate request ID
    if reqID, ok := ctx.Value(RequestIDKey).(string); ok {
        req.Header.Set("X-Request-ID", reqID)
    }
    
    client := &http.Client{Timeout: 15 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("service call failed: %w", err)
    }
    defer resp.Body.Close()
    
    var result Response
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }
    
    return &result, nil
}
```

---

## Key Takeaways

1. **Always use context** for cancellation, timeouts, and request-scoped data
2. **Make context the first parameter** in function signatures
3. **Always call defer cancel()** to prevent leaks
4. **Propagate context** through your call stack
5. **Check ctx.Done()** in long-running operations
6. **Use context values sparingly** - prefer explicit parameters
7. **Test cancellation behavior** to ensure responsiveness
8. **Chain contexts** to create proper hierarchies with appropriate timeouts

Context is essential for building responsive, well-behaved concurrent Go applications. Master these patterns and your applications will handle cancellation, timeouts, and coordination gracefully.