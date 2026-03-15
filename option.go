package fn

import "testing"

// Option represents an optional value of type A
type Option[A any] struct {
	isSome bool
	some   A
}

// Some wraps a value in an Option
func Some[A any](a A) Option[A] {
	return Option[A]{isSome: true, some: a}
}

// None returns an empty Option
func None[A any]() Option[A] {
	return Option[A]{}
}

// OptionFromPtr converts a pointer to an Option. nil becomes None
func OptionFromPtr[A any](p *A) Option[A] {
	if p == nil {
		return None[A]()
	}
	return Some(*p)
}

// IsSome returns true if the Option contains a value
func (o Option[A]) IsSome() bool {
	return o.isSome
}

// IsNone returns true if the Option is empty
func (o Option[A]) IsNone() bool {
	return !o.isSome
}

// UnwrapOr returns the contained value or the given default
func (o Option[A]) UnwrapOr(def A) A {
	if o.isSome {
		return o.some
	}
	return def
}

// UnwrapOrFunc returns the contained value or calls f to produce a default
func (o Option[A]) UnwrapOrFunc(f func() A) A {
	if o.isSome {
		return o.some
	}
	return f()
}

// UnwrapOrFail extracts the value or fails the test
func (o Option[A]) UnwrapOrFail(t *testing.T) A {
	t.Helper()
	if !o.isSome {
		t.Fatal("called UnwrapOrFail on None")
	}
	return o.some
}

// UnwrapOrErr returns the contained value or the provided error
func (o Option[A]) UnwrapOrErr(err error) (A, error) {
	if o.isSome {
		return o.some, nil
	}
	var zero A
	return zero, err
}

// WhenSome calls f with the contained value if present
func (o Option[A]) WhenSome(f func(A)) {
	if o.isSome {
		f(o.some)
	}
}

// Alt returns o if it is Some, otherwise returns other
func (o Option[A]) Alt(other Option[A]) Option[A] {
	if o.isSome {
		return o
	}
	return other
}

// UnsafeFromSome extracts the value, panicking if None
func (o Option[A]) UnsafeFromSome() A {
	if !o.isSome {
		panic("UnsafeFromSome called on None")
	}
	return o.some
}

// SomeToLeft converts Some(a) to Left(a), using the provided right for None
func (o Option[A]) SomeToLeft(right func() error) Either[A, error] {
	if o.isSome {
		return NewLeft[A, error](o.some)
	}
	return NewRight[A, error](right())
}

// SomeToRight converts Some(a) to Right(a), using the provided left for None
func (o Option[A]) SomeToRight(left func() error) Either[error, A] {
	if o.isSome {
		return NewRight[error, A](o.some)
	}
	return NewLeft[error, A](left())
}

// SomeToOk converts Some(a) to Ok(a), None to Err(err)
func (o Option[A]) SomeToOk(err error) Result[A] {
	if o.isSome {
		return Ok(o.some)
	}
	return Err[A](err)
}

// SomeToOkf converts Some(a) to Ok(a), None to Err with formatted message
func (o Option[A]) SomeToOkf(format string, args ...any) Result[A] {
	if o.isSome {
		return Ok(o.some)
	}
	return Errf[A](format, args...)
}

// ElimOption folds an Option: applies f to the value if Some, returns def if
// None
func ElimOption[A, B any](o Option[A], def B, f func(A) B) B {
	if o.isSome {
		return f(o.some)
	}
	return def
}

// MapOption applies f to the value inside an Option
func MapOption[A, B any](f func(A) B, o Option[A]) Option[B] {
	if o.isSome {
		return Some(f(o.some))
	}
	return None[B]()
}

// MapOptionZ is like MapOption but returns the zero value of B for None
func MapOptionZ[A, B any](f func(A) B, o Option[A]) B {
	if o.isSome {
		return f(o.some)
	}
	var zero B
	return zero
}

// FlatMapOption applies f to the value inside an Option, where f itself
// returns an Option
func FlatMapOption[A, B any](f func(A) Option[B], o Option[A]) Option[B] {
	if o.isSome {
		return f(o.some)
	}
	return None[B]()
}

// FlattenOption collapses a nested Option
func FlattenOption[A any](o Option[Option[A]]) Option[A] {
	if o.isSome {
		return o.some
	}
	return None[A]()
}

// LiftA2Option applies a binary function to two Options, returning None if
// either is None
func LiftA2Option[A, B, C any](
	f func(A, B) C, a Option[A], b Option[B],
) Option[C] {
	if a.isSome && b.isSome {
		return Some(f(a.some, b.some))
	}
	return None[C]()
}

// TransposeOptRes converts Option[Result[T]] to Result[Option[T]]
func TransposeOptRes[T any](o Option[Result[T]]) Result[Option[T]] {
	if !o.isSome {
		return Ok(None[T]())
	}

	val, err := o.some.Unpack()
	if err != nil {
		return Err[Option[T]](err)
	}

	return Ok(Some(val))
}
