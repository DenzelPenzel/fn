package fn

// ResultOpt represents a Result containing an optional value. It encodes the
// three states: success-with-value, success-without-value, and error
type ResultOpt[T any] struct {
	inner Result[Option[T]]
}

// OkOpt creates a ResultOpt with a present value
func OkOpt[T any](val T) ResultOpt[T] {
	return ResultOpt[T]{inner: Ok(Some(val))}
}

// NoneOpt creates a ResultOpt with no value (success but absent)
func NoneOpt[T any]() ResultOpt[T] {
	return ResultOpt[T]{inner: Ok(None[T]())}
}

// ErrOpt creates a ResultOpt with an error
func ErrOpt[T any](err error) ResultOpt[T] {
	return ResultOpt[T]{inner: Err[Option[T]](err)}
}

// IsSome returns true if the ResultOpt contains a value
func (ro ResultOpt[T]) IsSome() bool {
	if ro.inner.IsErr() {
		return false
	}
	return ro.inner.Either.left.IsSome()
}

// IsNone returns true if the ResultOpt is success-but-absent
func (ro ResultOpt[T]) IsNone() bool {
	if ro.inner.IsErr() {
		return false
	}
	return ro.inner.Either.left.IsNone()
}

// IsErr returns true if the ResultOpt contains an error
func (ro ResultOpt[T]) IsErr() bool {
	return ro.inner.IsErr()
}

// Unpack returns the inner Result[Option[T]]
func (ro ResultOpt[T]) Unpack() Result[Option[T]] {
	return ro.inner
}

// MapResultOpt applies f to the contained value if present
func MapResultOpt[A, B any](
	f func(A) B, ro ResultOpt[A],
) ResultOpt[B] {
	mapped := MapResultOk(
		func(opt Option[A]) Option[B] {
			return MapOption(f, opt)
		},
		ro.inner,
	)
	return ResultOpt[B]{inner: mapped}
}

// AndThenResultOpt applies f to the contained value if present, where f
// returns a ResultOpt
func AndThenResultOpt[A, B any](
	ro ResultOpt[A], f func(A) ResultOpt[B],
) ResultOpt[B] {
	if ro.inner.IsErr() {
		return ErrOpt[B](ro.inner.Either.right)
	}

	opt := ro.inner.Either.left
	if opt.IsNone() {
		return NoneOpt[B]()
	}

	return f(opt.some)
}
