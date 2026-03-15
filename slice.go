package fn

import "sync"

// Map applies f to each element of the slice, returning a new slice
func Map[A, B any](f func(A) B, xs []A) []B {
	result := make([]B, len(xs))
	for i, x := range xs {
		result[i] = f(x)
	}
	return result
}

// Filter returns elements of xs for which pred is true
func Filter[A any](pred Pred[A], xs []A) []A {
	var result []A
	for _, x := range xs {
		if pred(x) {
			result = append(result, x)
		}
	}
	return result
}

// FilterMap applies f to each element and collects the Some results
func FilterMap[A, B any](f func(A) Option[B], xs []A) []B {
	var result []B
	for _, x := range xs {
		opt := f(x)
		if opt.isSome {
			result = append(result, opt.some)
		}
	}
	return result
}

// TrimNones extracts values from a slice of Options, discarding Nones
func TrimNones[A any](xs []Option[A]) []A {
	var result []A
	for _, x := range xs {
		if x.isSome {
			result = append(result, x.some)
		}
	}
	return result
}

// Foldl is a left fold over a slice
func Foldl[A, B any](f func(B, A) B, init B, xs []A) B {
	acc := init
	for _, x := range xs {
		acc = f(acc, x)
	}
	return acc
}

// Foldr is a right fold over a slice
func Foldr[A, B any](f func(A, B) B, init B, xs []A) B {
	acc := init
	for i := len(xs) - 1; i >= 0; i-- {
		acc = f(xs[i], acc)
	}
	return acc
}

// Sum returns the sum of a numeric slice
func Sum[N Number](xs []N) N {
	var total N
	for _, x := range xs {
		total += x
	}
	return total
}

// Find returns the first element matching pred
func Find[A any](pred Pred[A], xs []A) Option[A] {
	for _, x := range xs {
		if pred(x) {
			return Some(x)
		}
	}
	return None[A]()
}

// FindIdx returns the first element matching pred along with its index
func FindIdx[A any](pred Pred[A], xs []A) Option[T2[int, A]] {
	for i, x := range xs {
		if pred(x) {
			return Some(NewT2(i, x))
		}
	}
	return None[T2[int, A]]()
}

// Elem returns true if the element is in the slice
func Elem[A comparable](a A, xs []A) bool {
	for _, x := range xs {
		if x == a {
			return true
		}
	}
	return false
}

// All returns true if pred is true for all elements
func All[A any](pred Pred[A], xs []A) bool {
	for _, x := range xs {
		if !pred(x) {
			return false
		}
	}
	return true
}

// Any returns true if pred is true for any element
func Any[A any](pred Pred[A], xs []A) bool {
	for _, x := range xs {
		if pred(x) {
			return true
		}
	}
	return false
}

// Flatten concatenates a slice of slices
func Flatten[A any](xss [][]A) []A {
	var result []A
	for _, xs := range xss {
		result = append(result, xs...)
	}
	return result
}

// Replicate creates a slice of n copies of a
func Replicate[A any](n int, a A) []A {
	result := make([]A, n)
	for i := range result {
		result[i] = a
	}
	return result
}

// Span splits xs into a prefix satisfying pred and the remainder
func Span[A any](pred Pred[A], xs []A) ([]A, []A) {
	for i, x := range xs {
		if !pred(x) {
			return xs[:i], xs[i:]
		}
	}
	return xs, nil
}

// SplitAt splits a slice at the given index
func SplitAt[A any](n int, xs []A) ([]A, []A) {
	if n < 0 {
		n = 0
	}
	if n > len(xs) {
		n = len(xs)
	}
	return xs[:n], xs[n:]
}

// ZipWith combines two slices element-wise using f, truncating to the shorter
// length
func ZipWith[A, B, C any](f func(A, B) C, as []A, bs []B) []C {
	n := len(as)
	if len(bs) < n {
		n = len(bs)
	}
	result := make([]C, n)
	for i := 0; i < n; i++ {
		result[i] = f(as[i], bs[i])
	}
	return result
}

// Head returns the first element of a slice
func Head[A any](xs []A) Option[A] {
	if len(xs) == 0 {
		return None[A]()
	}
	return Some(xs[0])
}

// Tail returns all elements after the first
func Tail[A any](xs []A) Option[[]A] {
	if len(xs) == 0 {
		return None[[]A]()
	}
	return Some(xs[1:])
}

// Init returns all elements except the last
func Init[A any](xs []A) Option[[]A] {
	if len(xs) == 0 {
		return None[[]A]()
	}
	return Some(xs[:len(xs)-1])
}

// Last returns the last element of a slice
func Last[A any](xs []A) Option[A] {
	if len(xs) == 0 {
		return None[A]()
	}
	return Some(xs[len(xs)-1])
}

// Uncons splits a slice into its head and tail
func Uncons[A any](xs []A) Option[T2[A, []A]] {
	if len(xs) == 0 {
		return None[T2[A, []A]]()
	}
	return Some(NewT2(xs[0], xs[1:]))
}

// Unsnoc splits a slice into its init and last
func Unsnoc[A any](xs []A) Option[T2[[]A, A]] {
	if len(xs) == 0 {
		return None[T2[[]A, A]]()
	}
	return Some(NewT2(xs[:len(xs)-1], xs[len(xs)-1]))
}

// SliceToMap converts a slice of key-value tuples to a map
func SliceToMap[K comparable, V any](xs []T2[K, V]) map[K]V {
	m := make(map[K]V, len(xs))
	for _, x := range xs {
		m[x.first] = x.second
	}
	return m
}

// HasDuplicates returns true if the slice contains duplicate elements
func HasDuplicates[A comparable](xs []A) bool {
	seen := make(map[A]struct{}, len(xs))
	for _, x := range xs {
		if _, ok := seen[x]; ok {
			return true
		}
		seen[x] = struct{}{}
	}
	return false
}

// Len returns the length of a slice
func Len[A any](xs []A) int {
	return len(xs)
}

// ForEachConc runs f on each element concurrently, waiting for all to finish
func ForEachConc[A any](f func(A), xs []A) {
	var wg sync.WaitGroup
	wg.Add(len(xs))
	for _, x := range xs {
		go func(val A) {
			defer wg.Done()
			f(val)
		}(x)
	}
	wg.Wait()
}

// CollectOptions collects all Some values from a slice of Options. If all are
// Some, returns Some(slice); otherwise returns None
func CollectOptions[A any](opts []Option[A]) Option[[]A] {
	result := make([]A, 0, len(opts))
	for _, o := range opts {
		if !o.isSome {
			return None[[]A]()
		}
		result = append(result, o.some)
	}
	return Some(result)
}

// CollectResults collects all Ok values from a slice of Results. If all are
// Ok, returns Ok(slice); otherwise returns the first error
func CollectResults[A any](rs []Result[A]) Result[[]A] {
	result := make([]A, 0, len(rs))
	for _, r := range rs {
		if r.IsErr() {
			return Err[[]A](r.Either.right)
		}
		result = append(result, r.Either.left)
	}
	return Ok(result)
}

// TraverseOption applies f to each element and collects results. Returns None
// if any application returns None
func TraverseOption[A, B any](f func(A) Option[B], xs []A) Option[[]B] {
	result := make([]B, 0, len(xs))
	for _, x := range xs {
		opt := f(x)
		if !opt.isSome {
			return None[[]B]()
		}
		result = append(result, opt.some)
	}
	return Some(result)
}

// TraverseResult applies f to each element and collects results. Returns the
// first error encountered
func TraverseResult[A, B any](f func(A) Result[B], xs []A) Result[[]B] {
	result := make([]B, 0, len(xs))
	for _, x := range xs {
		r := f(x)
		if r.IsErr() {
			return Err[[]B](r.Either.right)
		}
		result = append(result, r.Either.left)
	}
	return Ok(result)
}
