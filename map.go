package fn

// KeySet returns the set of keys from a map
func KeySet[K comparable, V any](m map[K]V) Set[K] {
	s := make(Set[K], len(m))
	for k := range m {
		s.Add(k)
	}
	return s
}

// NewSubMapIntersect returns a new map containing only the keys present in
// both the map and the set
func NewSubMapIntersect[K comparable, V any](
	m map[K]V, keys Set[K],
) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if keys.Contains(k) {
			result[k] = v
		}
	}
	return result
}

// NewSubMap returns a new map containing the values for each key in the set
// Missing keys are skipped
func NewSubMap[K comparable, V any](m map[K]V, keys Set[K]) map[K]V {
	result := make(map[K]V, keys.Size())
	for k := range keys {
		if v, ok := m[k]; ok {
			result[k] = v
		}
	}
	return result
}
