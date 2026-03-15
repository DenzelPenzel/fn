package fn

import "iter"

// Collect gathers all values from an iter.Seq into a slice
func Collect[T any](seq iter.Seq[T]) []T {
	var result []T
	for v := range seq {
		result = append(result, v)
	}
	return result
}

// CollectErr gathers all values from an iter.Seq2 of (T, error). It stops and
// returns the first error encountered
func CollectErr[T any](seq iter.Seq2[T, error]) ([]T, error) {
	var result []T
	for v, err := range seq {
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}
