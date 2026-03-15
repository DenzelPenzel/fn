// Package fn provides type-safe functional programming primitives for Go
// It has zero dependencies outside the standard library and relies on Go
// generics to deliver reusable, composable abstractions
//
// # Core Algebraic Data Types
//
// Option[A] models a value that may or may not be present, replacing
// nil-pointer conventions with explicit Some / None constructors
//
// Either[L, R] is a disjoint union of two types. By convention Left is the
// primary or success value, Right is the secondary or error value
//
// Result[T] specialises Either[T, error] for idiomatic Go error handling
// Construct with Ok, Err, or NewResult, then chain with MapOk, FlatMap,
// and friends
//
// ResultOpt[T] encodes three states: success-with-value, success-without-
// value, and error. Useful when absence is not an error
//
// T2[A, B] is a generic 2-tuple
//
// # Slice Operations
//
// Map, Filter, Foldl, Foldr, Find, ZipWith, Flatten, and many more operate
// on generic slices. CollectResults and TraverseResult give short-circuiting
// traversals over Result-returning functions
//
// # Collections
//
// Set[T] is a map-backed set with Union, Intersect, Diff, and Subset
// List[T] is a generic doubly-linked list. Queue[T] is a simple FIFO queue
//
// # Concurrency
//
// GoroutineManager manages goroutine lifecycles. ContextGuard ties spawned
// goroutines to a parent context. ConcurrentQueue[T] is an unbounded
// channel-bridging queue. EventDistributor broadcasts events to multiple
// subscribers. SendOrQuit, RecvOrTimeout, and Req provide channel helpers
//
// # Combinators
//
// Comp composes two functions. Iden is the identity function. Const returns
// a constant-valued function. Pred, PredAnd, PredOr, and PredNot combine
// predicates
package fn
