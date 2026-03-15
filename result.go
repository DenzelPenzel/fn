package fn

import (
	"fmt"
	"testing"
)

// Result represents a success value of type T or an error. It is implemented
// as an Either[T, error] where Left holds the success value
type Result[T any] struct {
	Either[T, error]
}

// Ok wraps a success value in a Result
func Ok[T any](val T) Result[T] {
	return Result[T]{Either: NewLeft[T, error](val)}
}

// Err wraps an error in a Result
func Err[T any](err error) Result[T] {
	return Result[T]{Either: NewRight[T, error](err)}
}

// Errf creates an error Result with a formatted message
func Errf[T any](format string, args ...any) Result[T] {
	return Err[T](fmt.Errorf(format, args...))
}

// NewResult creates a Result from a (value, error) pair
func NewResult[T any](val T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return Ok(val)
}

// Unpack returns the success value and error
func (r Result[T]) Unpack() (T, error) {
	if r.Either.isRight {
		var zero T
		return zero, r.Either.right
	}
	return r.Either.left, nil
}

// Err returns the error if present, nil otherwise
func (r Result[T]) Err() error {
	if r.Either.isRight {
		return r.Either.right
	}
	return nil
}

// IsOk returns true if the Result holds a success value
func (r Result[T]) IsOk() bool {
	return !r.Either.isRight
}

// IsErr returns true if the Result holds an error
func (r Result[T]) IsErr() bool {
	return r.Either.isRight
}

// MapOk applies f to the success value
func (r Result[T]) MapOk(f func(T) T) Result[T] {
	if r.IsOk() {
		return Ok(f(r.Either.left))
	}
	return r
}

// MapErr applies f to the error
func (r Result[T]) MapErr(f func(error) error) Result[T] {
	if r.IsErr() {
		return Err[T](f(r.Either.right))
	}
	return r
}

// OkToSome returns Some(val) if Ok, None if Err
func (r Result[T]) OkToSome() Option[T] {
	if r.IsOk() {
		return Some(r.Either.left)
	}
	return None[T]()
}

// WhenOk calls f with the success value if present
func (r Result[T]) WhenOk(f func(T)) {
	if r.IsOk() {
		f(r.Either.left)
	}
}

// WhenErr calls f with the error if present
func (r Result[T]) WhenErr(f func(error)) {
	if r.IsErr() {
		f(r.Either.right)
	}
}

// UnwrapOr returns the success value or the given default
func (r Result[T]) UnwrapOr(def T) T {
	if r.IsOk() {
		return r.Either.left
	}
	return def
}

// UnwrapOrElse returns the success value or calls f to produce a default
func (r Result[T]) UnwrapOrElse(f func() T) T {
	if r.IsOk() {
		return r.Either.left
	}
	return f()
}

// UnwrapOrFail extracts the success value or fails the test
func (r Result[T]) UnwrapOrFail(t *testing.T) T {
	t.Helper()
	if r.IsErr() {
		t.Fatalf("called UnwrapOrFail on Err: %v", r.Either.right)
	}
	return r.Either.left
}

// FlatMap applies f to the success value, where f returns a Result
func (r Result[T]) FlatMap(f func(T) Result[T]) Result[T] {
	if r.IsOk() {
		return f(r.Either.left)
	}
	return r
}

// AndThen is an alias for FlatMap
func (r Result[T]) AndThen(f func(T) Result[T]) Result[T] {
	return r.FlatMap(f)
}

// OrElse returns this Result if Ok, otherwise calls f
func (r Result[T]) OrElse(f func(error) Result[T]) Result[T] {
	if r.IsErr() {
		return f(r.Either.right)
	}
	return r
}

// Sink consumes the Result by calling okFn on success or errFn on error
func (r Result[T]) Sink(okFn func(T), errFn func(error)) {
	if r.IsOk() {
		okFn(r.Either.left)
	} else {
		errFn(r.Either.right)
	}
}

// MapResultOk applies a type-changing function to the success value. This is a
// free function because Go methods cannot introduce new type parameters
func MapResultOk[A, B any](f func(A) B, r Result[A]) Result[B] {
	if r.IsOk() {
		return Ok(f(r.Either.left))
	}
	return Err[B](r.Either.right)
}

// FlatMapResult applies f to the success value, where f returns a Result of a
// different type
func FlatMapResult[A, B any](f func(A) Result[B], r Result[A]) Result[B] {
	if r.IsOk() {
		return f(r.Either.left)
	}
	return Err[B](r.Either.right)
}

// FlattenResult collapses a nested Result
func FlattenResult[T any](r Result[Result[T]]) Result[T] {
	if r.IsOk() {
		return r.Either.left
	}
	return Err[T](r.Either.right)
}

// AndThenResult is a free-function variant of AndThen for type-changing
// operations
func AndThenResult[A, B any](r Result[A], f func(A) Result[B]) Result[B] {
	return FlatMapResult(f, r)
}

// LiftA2Result applies a binary function to two Results, returning the first
// error encountered
func LiftA2Result[A, B, C any](
	f func(A, B) C, a Result[A], b Result[B],
) Result[C] {
	if a.IsErr() {
		return Err[C](a.Either.right)
	}
	if b.IsErr() {
		return Err[C](b.Either.right)
	}
	return Ok(f(a.Either.left, b.Either.left))
}

// TransposeResOpt converts Result[Option[T]] to Option[Result[T]]
func TransposeResOpt[T any](r Result[Option[T]]) Option[Result[T]] {
	if r.IsErr() {
		return Some(Err[T](r.Either.right))
	}

	opt := r.Either.left
	if opt.IsNone() {
		return None[Result[T]]()
	}

	return Some(Ok(opt.some))
}
