package fn

// Set is a collection of unique elements backed by a map
type Set[T comparable] map[T]struct{}

// NewSet creates a Set from the given elements
func NewSet[T comparable](elems ...T) Set[T] {
	s := make(Set[T], len(elems))
	for _, e := range elems {
		s[e] = struct{}{}
	}
	return s
}

// Add inserts an element into the set
func (s Set[T]) Add(elem T) {
	s[elem] = struct{}{}
}

// Remove deletes an element from the set
func (s Set[T]) Remove(elem T) {
	delete(s, elem)
}

// Contains returns true if the element is in the set
func (s Set[T]) Contains(elem T) bool {
	_, ok := s[elem]
	return ok
}

// IsEmpty returns true if the set has no elements
func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

// Size returns the number of elements in the set
func (s Set[T]) Size() int {
	return len(s)
}

// Diff returns elements in s that are not in other
func (s Set[T]) Diff(other Set[T]) Set[T] {
	result := NewSet[T]()
	for elem := range s {
		if !other.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Union returns elements in either set
func (s Set[T]) Union(other Set[T]) Set[T] {
	result := s.Copy()
	for elem := range other {
		result.Add(elem)
	}
	return result
}

// Intersect returns elements in both sets
func (s Set[T]) Intersect(other Set[T]) Set[T] {
	result := NewSet[T]()
	for elem := range s {
		if other.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Subset returns true if all elements of s are in other
func (s Set[T]) Subset(other Set[T]) bool {
	for elem := range s {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// Equal returns true if both sets contain the same elements
func (s Set[T]) Equal(other Set[T]) bool {
	return len(s) == len(other) && s.Subset(other)
}

// ToSlice returns the elements as a slice (order not guaranteed)
func (s Set[T]) ToSlice() []T {
	result := make([]T, 0, len(s))
	for elem := range s {
		result = append(result, elem)
	}
	return result
}

// Copy returns a shallow copy of the set
func (s Set[T]) Copy() Set[T] {
	result := make(Set[T], len(s))
	for elem := range s {
		result[elem] = struct{}{}
	}
	return result
}

// SetDiff is a free-function variant of Diff
func SetDiff[T comparable](a, b Set[T]) Set[T] {
	return a.Diff(b)
}
