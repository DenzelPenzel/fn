package fn

import "testing"

func TestKeySet(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	ks := KeySet(m)
	assertEqual(t, ks.Size(), 3)
	assertTrue(t, ks.Contains("a"))
	assertTrue(t, ks.Contains("b"))
	assertTrue(t, ks.Contains("c"))
}

func TestNewSubMapIntersect(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := NewSet("a", "c", "d")
	sub := NewSubMapIntersect(m, keys)
	assertEqual(t, len(sub), 2)
	assertEqual(t, sub["a"], 1)
	assertEqual(t, sub["c"], 3)
}

func TestNewSubMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := NewSet("a", "c", "d")
	sub := NewSubMap(m, keys)
	assertEqual(t, len(sub), 2)
	assertEqual(t, sub["a"], 1)
	assertEqual(t, sub["c"], 3)
}
