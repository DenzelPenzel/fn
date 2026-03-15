# fn

[![Go Reference](https://pkg.go.dev/badge/github.com/denzelpenzel/fn.svg)](https://pkg.go.dev/github.com/denzelpenzel/fn)
[![CI](https://github.com/denzelpenzel/fn/actions/workflows/ci.yml/badge.svg)](https://github.com/denzelpenzel/fn/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Type-safe functional programming primitives for Go.

## Install

```bash
go get github.com/denzelpenzel/fn
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/denzelpenzel/fn"
)

func main() {
    // Option — explicit presence/absence
    name := fn.Some("Alice")
    fmt.Println(name.UnwrapOr("unknown")) // Alice

    // Result — success or error
    r := fn.NewResult(42, nil)
    fmt.Println(r.IsOk()) // true

    // Slice ops — generic map, filter, fold
    doubled := fn.Map(func(x int) int { return x * 2 }, []int{1, 2, 3})
    fmt.Println(doubled) // [2 4 6]

    // Set — union, intersect, diff
    s := fn.NewSet(1, 2, 3).Union(fn.NewSet(3, 4))
    fmt.Println(s.Contains(4)) // true
}
```

## API Overview

| Category | Key Types / Functions |
|---|---|
| **Core** | `Unit`, `Eq`, `Neq`, `Comp`, `Iden`, `Const` |
| **Tuple** | `T2`, `NewT2`, `Pair`, `MapFirst`, `MapSecond` |
| **Option** | `Option`, `Some`, `None`, `MapOption`, `FlatMapOption`, `ElimOption`, `LiftA2Option` |
| **Either** | `Either`, `NewLeft`, `NewRight`, `ElimEither`, `MapLeft`, `MapRight` |
| **Result** | `Result`, `Ok`, `Err`, `NewResult`, `MapResultOk`, `FlatMapResult`, `LiftA2Result` |
| **ResultOpt** | `ResultOpt`, `OkOpt`, `NoneOpt`, `ErrOpt` |
| **Predicate** | `Pred`, `PredAnd`, `PredOr`, `PredNot` |
| **Slice** | `Map`, `Filter`, `Foldl`, `Foldr`, `Find`, `ZipWith`, `Flatten`, `Sum`, `All`, `Any`, `Head`, `Tail`, `Span`, `SplitAt`, `CollectResults`, `TraverseResult` |
| **Set** | `Set`, `NewSet`, `Add`, `Remove`, `Contains`, `Union`, `Intersect`, `Diff`, `Subset` |
| **Map** | `KeySet`, `NewSubMap`, `NewSubMapIntersect` |
| **List** | `List`, `NewList`, `PushBack`, `PushFront`, `Remove`, `Front`, `Back` |
| **Queue** | `Queue`, `NewQueue`, `Enqueue`, `Dequeue`, `Peek` |
| **Iterator** | `Collect`, `CollectErr` |
| **Channel** | `SendOrQuit`, `RecvOrTimeout`, `RecvResp`, `Req`, `NewReq` |
| **Concurrency** | `GoroutineManager`, `ContextGuard`, `ConcurrentQueue`, `EventDistributor`, `EventReceiver` |
| **I/O** | `WriteFile`, `WriteFileRemove` |
| **Testing** | `GuardTest` |

## Requirements

- Go 1.25+
- Zero dependencies

## Contributing

1. Fork the repository
2. Create a feature branch
3. `make test` — all tests must pass
4. `make lint` — code must be lint-clean
5. Open a pull request

## License

[MIT](LICENSE)
