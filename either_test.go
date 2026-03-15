package fn

import "testing"

func TestEitherLeft(t *testing.T) {
	e := NewLeft[int, string](42)
	assertTrue(t, e.IsLeft())
	assertFalse(t, e.IsRight())
	assertEqual(t, e.UnwrapLeftOr(0), 42)
	assertEqual(t, e.UnwrapRightOr("default"), "default")
}

func TestEitherRight(t *testing.T) {
	e := NewRight[int, string]("hello")
	assertFalse(t, e.IsLeft())
	assertTrue(t, e.IsRight())
	assertEqual(t, e.UnwrapLeftOr(0), 0)
	assertEqual(t, e.UnwrapRightOr("default"), "hello")
}

func TestEitherWhenLeft(t *testing.T) {
	var called bool
	e := NewLeft[int, string](5)
	e.WhenLeft(func(v int) { called = true; assertEqual(t, v, 5) })
	assertTrue(t, called)

	called = false
	e2 := NewRight[int, string]("x")
	e2.WhenLeft(func(_ int) { called = true })
	assertFalse(t, called)
}

func TestEitherWhenRight(t *testing.T) {
	var called bool
	e := NewRight[int, string]("hi")
	e.WhenRight(func(v string) { called = true; assertEqual(t, v, "hi") })
	assertTrue(t, called)
}

func TestEitherLeftToSome(t *testing.T) {
	e := NewLeft[int, string](10)
	opt := e.LeftToSome()
	assertTrue(t, opt.IsSome())
	assertEqual(t, opt.UnsafeFromSome(), 10)

	e2 := NewRight[int, string]("x")
	assertTrue(t, e2.LeftToSome().IsNone())
}

func TestEitherRightToSome(t *testing.T) {
	e := NewRight[int, string]("ok")
	opt := e.RightToSome()
	assertTrue(t, opt.IsSome())
	assertEqual(t, opt.UnsafeFromSome(), "ok")

	e2 := NewLeft[int, string](1)
	assertTrue(t, e2.RightToSome().IsNone())
}

func TestEitherSwap(t *testing.T) {
	e := NewLeft[int, string](42)
	swapped := e.Swap()
	assertTrue(t, swapped.IsRight())
	assertEqual(t, swapped.UnwrapRightOr(0), 42)
}

func TestElimEither(t *testing.T) {
	left := NewLeft[int, string](5)
	result := ElimEither(left, func(_ int) string {
		return "left"
	}, func(_ string) string {
		return "right"
	})
	assertEqual(t, result, "left")

	right := NewRight[int, string]("hello")
	result2 := ElimEither(right, func(_ int) string {
		return "left"
	}, func(s string) string {
		return s
	})
	assertEqual(t, result2, "hello")
}

func TestMapLeft(t *testing.T) {
	e := NewLeft[int, string](5)
	mapped := MapLeft(func(i int) int { return i * 2 }, e)
	assertEqual(t, mapped.UnwrapLeftOr(0), 10)

	e2 := NewRight[int, string]("err")
	mapped2 := MapLeft(func(i int) int { return i * 2 }, e2)
	assertEqual(t, mapped2.UnwrapRightOr(""), "err")
}

func TestMapRight(t *testing.T) {
	e := NewRight[int, string]("hi")
	mapped := MapRight(func(s string) string { return s + "!" }, e)
	assertEqual(t, mapped.UnwrapRightOr(""), "hi!")

	e2 := NewLeft[int, string](1)
	mapped2 := MapRight(func(s string) string { return s + "!" }, e2)
	assertEqual(t, mapped2.UnwrapLeftOr(0), 1)
}
