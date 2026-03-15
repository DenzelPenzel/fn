package fn

import "testing"

func TestPredAnd(t *testing.T) {
	positive := func(x int) bool { return x > 0 }
	even := func(x int) bool { return x%2 == 0 }

	both := PredAnd(positive, even)
	assertTrue(t, both(4))
	assertFalse(t, both(3))
	assertFalse(t, both(-2))

	// Empty and: vacuously true
	assertTrue(t, PredAnd[int]()(0))
}

func TestPredOr(t *testing.T) {
	positive := func(x int) bool { return x > 0 }
	even := func(x int) bool { return x%2 == 0 }

	either := PredOr(positive, even)
	assertTrue(t, either(3))
	assertTrue(t, either(-2))
	assertFalse(t, either(-3))

	// Empty or: vacuously false
	assertFalse(t, PredOr[int]()(0))
}

func TestPredNot(t *testing.T) {
	positive := func(x int) bool { return x > 0 }
	notPositive := PredNot(positive)

	assertTrue(t, notPositive(-1))
	assertFalse(t, notPositive(1))
}
