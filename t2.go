package fn

// T2 is a 2-tuple holding values of types A and B
type T2[A, B any] struct {
	first  A
	second B
}

// NewT2 creates a new 2-tuple
func NewT2[A, B any](a A, b B) T2[A, B] {
	return T2[A, B]{first: a, second: b}
}

// Pair is an alias for NewT2
func Pair[A, B any](a A, b B) T2[A, B] {
	return NewT2(a, b)
}

// First returns the first element of the tuple
func (t T2[A, B]) First() A {
	return t.first
}

// Second returns the second element of the tuple
func (t T2[A, B]) Second() B {
	return t.second
}

// Unpack returns both elements of the tuple
func (t T2[A, B]) Unpack() (A, B) {
	return t.first, t.second
}

// MapFirst applies f to the first element
func MapFirst[A, B, C any](f func(A) C, t T2[A, B]) T2[C, B] {
	return NewT2(f(t.first), t.second)
}

// MapSecond applies f to the second element
func MapSecond[A, B, C any](f func(B) C, t T2[A, B]) T2[A, C] {
	return NewT2(t.first, f(t.second))
}
