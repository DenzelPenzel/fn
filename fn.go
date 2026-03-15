package fn

// Unit is the empty type with a single value
type Unit = struct{}

// Number is a constraint for all numeric types
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// Copyable is a constraint for types that can be copied without error
type Copyable[T any] interface {
	Copy() T
}

// CopyableErr is a constraint for types whose copy may fail
type CopyableErr[T any] interface {
	Copy() (T, error)
}

// Comp composes two functions: Comp(f, g)(x) == f(g(x))
func Comp[A, B, C any](f func(B) C, g func(A) B) func(A) C {
	return func(a A) C {
		return f(g(a))
	}
}

// Iden is the identity function
func Iden[A any](a A) A {
	return a
}

// Const returns a function that always returns the given value
func Const[A, B any](a A) func(B) A {
	return func(_ B) A {
		return a
	}
}

// Eq returns true if the two values are equal
func Eq[A comparable](a, b A) bool {
	return a == b
}

// Neq returns true if the two values are not equal
func Neq[A comparable](a, b A) bool {
	return a != b
}

// CopyAll copies a slice of Copyable values
func CopyAll[T any, S Copyable[T]](xs []S) []T {
	result := make([]T, len(xs))
	for i, x := range xs {
		result[i] = x.Copy()
	}
	return result
}

// CopyAllErr copies a slice of CopyableErr values, returning on first error
func CopyAllErr[T any, S CopyableErr[T]](xs []S) ([]T, error) {
	result := make([]T, len(xs))
	for i, x := range xs {
		c, err := x.Copy()
		if err != nil {
			return nil, err
		}
		result[i] = c
	}
	return result, nil
}
