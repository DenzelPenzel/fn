package fn

// Either represents a disjoint union of types L and R. By convention, Left is
// the "primary" or "success" value and Right is the "secondary" or "error"
// value
type Either[L, R any] struct {
	isRight bool
	left    L
	right   R
}

// NewLeft constructs an Either holding a Left value
func NewLeft[L, R any](l L) Either[L, R] {
	return Either[L, R]{left: l}
}

// NewRight constructs an Either holding a Right value
func NewRight[L, R any](r R) Either[L, R] {
	return Either[L, R]{isRight: true, right: r}
}

// IsLeft returns true if this Either holds a Left value
func (e Either[L, R]) IsLeft() bool {
	return !e.isRight
}

// IsRight returns true if this Either holds a Right value
func (e Either[L, R]) IsRight() bool {
	return e.isRight
}

// WhenLeft calls f with the Left value if present
func (e Either[L, R]) WhenLeft(f func(L)) {
	if !e.isRight {
		f(e.left)
	}
}

// WhenRight calls f with the Right value if present
func (e Either[L, R]) WhenRight(f func(R)) {
	if e.isRight {
		f(e.right)
	}
}

// LeftToSome returns Some(left) if Left, None otherwise
func (e Either[L, R]) LeftToSome() Option[L] {
	if !e.isRight {
		return Some(e.left)
	}
	return None[L]()
}

// RightToSome returns Some(right) if Right, None otherwise
func (e Either[L, R]) RightToSome() Option[R] {
	if e.isRight {
		return Some(e.right)
	}
	return None[R]()
}

// UnwrapLeftOr returns the Left value or the given default
func (e Either[L, R]) UnwrapLeftOr(def L) L {
	if !e.isRight {
		return e.left
	}
	return def
}

// UnwrapRightOr returns the Right value or the given default
func (e Either[L, R]) UnwrapRightOr(def R) R {
	if e.isRight {
		return e.right
	}
	return def
}

// Swap exchanges Left and Right
func (e Either[L, R]) Swap() Either[R, L] {
	if e.isRight {
		return NewLeft[R, L](e.right)
	}
	return NewRight[R, L](e.left)
}

// ElimEither folds an Either: applies f to Left or g to Right
func ElimEither[L, R, O any](e Either[L, R], f func(L) O, g func(R) O) O {
	if !e.isRight {
		return f(e.left)
	}
	return g(e.right)
}

// MapLeft transforms the Left value of an Either
func MapLeft[L, R, O any](f func(L) O, e Either[L, R]) Either[O, R] {
	if !e.isRight {
		return NewLeft[O, R](f(e.left))
	}
	return NewRight[O, R](e.right)
}

// MapRight transforms the Right value of an Either
func MapRight[L, R, O any](f func(R) O, e Either[L, R]) Either[L, O] {
	if !e.isRight {
		return NewLeft[L, O](e.left)
	}
	return NewRight[L, O](f(e.right))
}
