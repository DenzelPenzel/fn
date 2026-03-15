package fn

import "testing"

func TestT2(t *testing.T) {
	p := NewT2(1, "hello")
	assertEqual(t, p.First(), 1)
	assertEqual(t, p.Second(), "hello")

	a, b := p.Unpack()
	assertEqual(t, a, 1)
	assertEqual(t, b, "hello")
}

func TestPair(t *testing.T) {
	p := Pair(true, 3.14)
	assertEqual(t, p.First(), true)
	assertEqual(t, p.Second(), 3.14)
}

func TestMapFirst(t *testing.T) {
	p := NewT2(3, "x")
	result := MapFirst(func(i int) int { return i * 2 }, p)
	assertEqual(t, result.First(), 6)
	assertEqual(t, result.Second(), "x")
}

func TestMapSecond(t *testing.T) {
	p := NewT2("x", 3)
	result := MapSecond(func(i int) string { return "y" }, p)
	assertEqual(t, result.First(), "x")
	assertEqual(t, result.Second(), "y")
}
