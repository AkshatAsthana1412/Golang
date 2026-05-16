# Go Concepts — Senior Interview Practice

A collection of ~80 hands-on problems covering Go beyond concurrency: types, interfaces, OOP idioms, generics, errors, memory model, the standard library gotchas, and the Gin web framework. Designed for engineers preparing for senior-level Go interviews.

## How to Use

Each problem lives in its own directory with two files:

- **`main.go`** — Problem statement, hints, and starter scaffold. **Edit this.**
- **`solution/main.go`** — Reference implementation with explanatory comments. Read *after* you attempt.

Run a problem from its directory:

```bash
cd go-concepts/01-types-and-methods/01-zero-values
go run .            # your attempt
go run ./solution   # reference
```

---

## Problem Index

### 01 — Types & Methods
| # | Problem | Concepts |
|---|---------|----------|
| 1 | Zero Values | default values, declared vs initialized |
| 2 | Type Definitions vs Aliases | `type T = U` vs `type T U` |
| 3 | Method Sets | what methods a value/pointer has |
| 4 | Value vs Pointer Receivers | when each is correct |
| 5 | Named Types & Underlying Types | conversion rules |
| 6 | Struct Comparability | when `==` works on structs |
| 7 | Anonymous Structs | one-off types, JSON shapes |
| 8 | Function Types as First-Class Values | callbacks, method values |

### 02 — Interfaces & OOP (Encapsulation, Abstraction)
| # | Problem | Concepts |
|---|---------|----------|
| 1 | Implicit Satisfaction | duck typing, no `implements` keyword |
| 2 | Empty Interface (`any`) | type-erased containers |
| 3 | Type Assertions | `v, ok := x.(T)` |
| 4 | Type Switches | `switch v := x.(type)` |
| 5 | Encapsulation via Unexported Fields | package-private state, exported API |
| 6 | Abstraction via Interface | depend on behavior, not types |
| 7 | Interface Composition | small interfaces combined |
| 8 | The Nil Interface Trap | typed nil != nil interface |
| 9 | Accept Interfaces, Return Structs | API design proverb |
| 10 | `io.Reader` / `io.Writer` Pipeline | the canonical Go interface |
| 11 | `sort.Interface` | implementing standard hooks |
| 12 | `fmt.Stringer` & Custom Formatting | participating in fmt |

### 03 — Embedding & Composition
| # | Problem | Concepts |
|---|---------|----------|
| 1 | Struct Embedding | composition over inheritance |
| 2 | Method Promotion | how embedded methods surface |
| 3 | Interface Embedding | building bigger interfaces |
| 4 | Ambiguity & Shadowing | when promoted names collide |
| 5 | Mixins via Embedding | adding capabilities to types |

### 04 — Error Handling & Wrapping
| # | Problem | Concepts |
|---|---------|----------|
| 1 | `error` Interface | nothing magical |
| 2 | Sentinel Errors | `var ErrNotFound = errors.New(...)` |
| 3 | Custom Error Types | structured error data |
| 4 | Wrapping with `%w` | `fmt.Errorf("...: %w", err)` |
| 5 | `errors.Is` & `errors.As` | unwrapping chains |
| 6 | Panic vs Error | when to panic |
| 7 | Recover in Defer | converting panics into errors |
| 8 | Multi-Errors with `errors.Join` | aggregating failures |

### 05 — Generics
| # | Problem | Concepts |
|---|---------|----------|
| 1 | Generic Function: `Map` | `[T, U any]` type parameters |
| 2 | Generic Function: `Filter` | predicate over `[]T` |
| 3 | `comparable` Constraint | sets, dedup |
| 4 | `cmp.Ordered` Constraint | min/max/sort |
| 5 | Custom Constraint Interface | type sets with `~` |
| 6 | Generic Stack | parameterized data structure |
| 7 | Generic LRU-ish Cache | parameterized state |
| 8 | Type Inference Limits | when you must specify |

### 06 — Memory Model & Escape Analysis
| # | Problem | Concepts |
|---|---------|----------|
| 1 | Stack vs Heap | where does this allocate? |
| 2 | Forcing Escape | returning a pointer to a local |
| 3 | Interface Boxing | escape via `any` |
| 4 | Slice Backing Array Escape | append behavior |
| 5 | Struct Padding & Alignment | field ordering matters |
| 6 | Pointer vs Value: Cost | when copying is cheaper |
| 7 | Detecting Escapes | `-gcflags="-m"` |

### 07 — Slices, Maps, Arrays
| # | Problem | Concepts |
|---|---------|----------|
| 1 | Slice Header Anatomy | ptr/len/cap |
| 2 | `append` Aliasing Bug | shared backing arrays |
| 3 | Slice as Function Argument | mutating vs reslicing |
| 4 | `copy` Semantics | min-of-lens, no growth |
| 5 | Map Iteration Order | randomized on purpose |
| 6 | Nil Map Read vs Write | one is fine, one panics |
| 7 | Deletion During Iteration | safe in Go |
| 8 | Comparing Slices | `reflect.DeepEqual` / `slices.Equal` |

### 08 — Strings, Bytes, Runes
| # | Problem | Concepts |
|---|---------|----------|
| 1 | UTF-8 Indexing Pitfall | `s[i]` returns a byte |
| 2 | Range over String | rune-by-rune |
| 3 | `[]byte` ↔ `string` Conversions | allocation cost |
| 4 | `strings.Builder` | efficient concatenation |
| 5 | Byte/Rune Counting | `len` vs `utf8.RuneCountInString` |

### 09 — Gin Web Framework
| # | Problem | Concepts |
|---|---------|----------|
| 1 | Hello Server | `gin.Default()`, routes |
| 2 | Path & Query Params | `c.Param`, `c.Query` |
| 3 | JSON Binding | `c.ShouldBindJSON` + structs |
| 4 | Validator Tags | `binding:"required,email"` |
| 5 | Custom Validator | `binding.Validator` registration |
| 6 | Middleware Basics | `gin.HandlerFunc`, `c.Next` |
| 7 | Auth Middleware + `c.Set/Get` | request-scoped values |
| 8 | Route Groups | `r.Group("/api/v1")` |
| 9 | Centralized Error Handling | `c.Error` + middleware |
| 10 | Request Cancellation | `c.Request.Context()` |
| 11 | File Upload | multipart forms |
| 12 | Graceful Shutdown | `http.Server.Shutdown` with Gin |

### 10 — Misc Must-Know
| # | Problem | Concepts |
|---|---------|----------|
| 1 | Defer Order & Loop Pitfall | LIFO, captured values |
| 2 | `iota` Patterns | enums, bit flags |
| 3 | Struct Tags & Reflection | `json:"name,omitempty"` |
| 4 | Custom JSON Marshal/Unmarshal | `MarshalJSON` / `UnmarshalJSON` |
| 5 | `init` Functions | order, side-effects |
| 6 | `context.Context` Plumbing | deadlines, values, cancellation |
| 7 | Build Tags & Conditional Compilation | `//go:build` |

---

## Conventions

- Every `main.go` compiles as-is (with TODOs left blank). Implement the TODOs and run.
- Solutions favor clarity over cleverness; each major step is commented.
- Topics 1–8 and 10 use only the standard library. Topic 9 (Gin) requires `go get github.com/gin-gonic/gin`.
