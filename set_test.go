package fn

import "testing"

func TestSetBasic(t *testing.T) {
	s := NewSet(1, 2, 3)
	assertEqual(t, s.Size(), 3)
	assertTrue(t, s.Contains(1))
	assertFalse(t, s.Contains(4))
	assertFalse(t, s.IsEmpty())

	s.Remove(2)
	assertEqual(t, s.Size(), 2)
	assertFalse(t, s.Contains(2))

	s.Add(4)
	assertTrue(t, s.Contains(4))
}

func TestSetEmpty(t *testing.T) {
	s := NewSet[int]()
	assertTrue(t, s.IsEmpty())
	assertEqual(t, s.Size(), 0)
}

func TestSetDiff(t *testing.T) {
	a := NewSet(1, 2, 3, 4)
	b := NewSet(2, 4, 6)
	diff := a.Diff(b)
	assertTrue(t, diff.Contains(1))
	assertTrue(t, diff.Contains(3))
	assertFalse(t, diff.Contains(2))
	assertEqual(t, diff.Size(), 2)

	// Free function variant
	diff2 := SetDiff(a, b)
	assertTrue(t, diff.Equal(diff2))
}

func TestSetUnion(t *testing.T) {
	a := NewSet(1, 2)
	b := NewSet(2, 3)
	u := a.Union(b)
	assertEqual(t, u.Size(), 3)
	assertTrue(t, u.Contains(1))
	assertTrue(t, u.Contains(2))
	assertTrue(t, u.Contains(3))
}

func TestSetIntersect(t *testing.T) {
	a := NewSet(1, 2, 3)
	b := NewSet(2, 3, 4)
	inter := a.Intersect(b)
	assertEqual(t, inter.Size(), 2)
	assertTrue(t, inter.Contains(2))
	assertTrue(t, inter.Contains(3))
}

func TestSetSubset(t *testing.T) {
	a := NewSet(1, 2)
	b := NewSet(1, 2, 3)
	assertTrue(t, a.Subset(b))
	assertFalse(t, b.Subset(a))
}

func TestSetEqual(t *testing.T) {
	a := NewSet(1, 2, 3)
	b := NewSet(3, 2, 1)
	assertTrue(t, a.Equal(b))

	c := NewSet(1, 2)
	assertFalse(t, a.Equal(c))
}

func TestSetCopy(t *testing.T) {
	a := NewSet(1, 2, 3)
	b := a.Copy()
	assertTrue(t, a.Equal(b))

	// Mutating copy should not affect original
	b.Add(4)
	assertFalse(t, a.Contains(4))
}

func TestSetToSlice(t *testing.T) {
	s := NewSet(1, 2, 3)
	sl := s.ToSlice()
	assertEqual(t, len(sl), 3)

	// All elements present
	for _, v := range sl {
		assertTrue(t, s.Contains(v))
	}
}
