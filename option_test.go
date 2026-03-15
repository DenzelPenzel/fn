package fn

import (
	"errors"
	"testing"
)

func TestOptionSome(t *testing.T) {
	o := Some(42)
	assertTrue(t, o.IsSome())
	assertFalse(t, o.IsNone())
	assertEqual(t, o.UnwrapOr(0), 42)
}

func TestOptionNone(t *testing.T) {
	o := None[int]()
	assertFalse(t, o.IsSome())
	assertTrue(t, o.IsNone())
	assertEqual(t, o.UnwrapOr(99), 99)
}

func TestOptionFromPtr(t *testing.T) {
	v := 42
	o := OptionFromPtr(&v)
	assertTrue(t, o.IsSome())
	assertEqual(t, o.UnsafeFromSome(), 42)

	o2 := OptionFromPtr[int](nil)
	assertTrue(t, o2.IsNone())
}

func TestOptionUnwrapOrFunc(t *testing.T) {
	o := None[int]()
	assertEqual(t, o.UnwrapOrFunc(func() int { return 7 }), 7)

	o2 := Some(3)
	assertEqual(t, o2.UnwrapOrFunc(func() int { return 7 }), 3)
}

func TestOptionUnwrapOrErr(t *testing.T) {
	sentinel := errors.New("missing")

	o := Some(5)
	val, err := o.UnwrapOrErr(sentinel)
	assertNoError(t, err)
	assertEqual(t, val, 5)

	o2 := None[int]()
	_, err = o2.UnwrapOrErr(sentinel)
	assertTrue(t, errors.Is(err, sentinel))
}

func TestOptionWhenSome(t *testing.T) {
	var called bool
	Some(10).WhenSome(func(v int) { called = true; assertEqual(t, v, 10) })
	assertTrue(t, called)

	called = false
	None[int]().WhenSome(func(_ int) { called = true })
	assertFalse(t, called)
}

func TestOptionAlt(t *testing.T) {
	a := None[int]()
	b := Some(5)
	assertEqual(t, a.Alt(b).UnsafeFromSome(), 5)

	c := Some(3)
	assertEqual(t, c.Alt(b).UnsafeFromSome(), 3)
}

func TestOptionUnsafeFromSomePanic(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic")
		}
	}()
	None[int]().UnsafeFromSome()
}

func TestOptionSomeToLeft(t *testing.T) {
	o := Some(5)
	e := o.SomeToLeft(func() error { return errors.New("err") })
	assertTrue(t, e.IsLeft())
	assertEqual(t, e.UnwrapLeftOr(0), 5)

	o2 := None[int]()
	e2 := o2.SomeToLeft(func() error { return errors.New("err") })
	assertTrue(t, e2.IsRight())
}

func TestOptionSomeToRight(t *testing.T) {
	o := Some(5)
	e := o.SomeToRight(func() error { return errors.New("err") })
	assertTrue(t, e.IsRight())

	o2 := None[int]()
	e2 := o2.SomeToRight(func() error { return errors.New("err") })
	assertTrue(t, e2.IsLeft())
}

func TestOptionSomeToOk(t *testing.T) {
	o := Some(42)
	r := o.SomeToOk(errors.New("missing"))
	assertTrue(t, r.IsOk())
	assertEqual(t, r.UnwrapOr(0), 42)

	o2 := None[int]()
	r2 := o2.SomeToOk(errors.New("missing"))
	assertTrue(t, r2.IsErr())
}

func TestElimOption(t *testing.T) {
	result := ElimOption(Some(3), "none", func(i int) string { return "some" })
	assertEqual(t, result, "some")

	result2 := ElimOption(None[int](), "none", func(i int) string { return "some" })
	assertEqual(t, result2, "none")
}

func TestMapOption(t *testing.T) {
	mapped := MapOption(func(i int) string { return "ok" }, Some(1))
	assertTrue(t, mapped.IsSome())
	assertEqual(t, mapped.UnsafeFromSome(), "ok")

	mapped2 := MapOption(func(i int) string { return "ok" }, None[int]())
	assertTrue(t, mapped2.IsNone())
}

func TestMapOptionZ(t *testing.T) {
	result := MapOptionZ(func(i int) int { return i * 2 }, Some(3))
	assertEqual(t, result, 6)

	result2 := MapOptionZ(func(i int) int { return i * 2 }, None[int]())
	assertEqual(t, result2, 0)
}

func TestFlatMapOption(t *testing.T) {
	f := func(i int) Option[string] {
		if i > 0 {
			return Some("positive")
		}
		return None[string]()
	}

	result := FlatMapOption(f, Some(1))
	assertEqual(t, result.UnsafeFromSome(), "positive")

	result2 := FlatMapOption(f, Some(-1))
	assertTrue(t, result2.IsNone())

	result3 := FlatMapOption(f, None[int]())
	assertTrue(t, result3.IsNone())
}

func TestFlattenOption(t *testing.T) {
	nested := Some(Some(42))
	assertEqual(t, FlattenOption(nested).UnsafeFromSome(), 42)

	nested2 := Some(None[int]())
	assertTrue(t, FlattenOption(nested2).IsNone())

	nested3 := None[Option[int]]()
	assertTrue(t, FlattenOption(nested3).IsNone())
}

func TestLiftA2Option(t *testing.T) {
	add := func(a, b int) int { return a + b }
	result := LiftA2Option(add, Some(3), Some(4))
	assertEqual(t, result.UnsafeFromSome(), 7)

	result2 := LiftA2Option(add, None[int](), Some(4))
	assertTrue(t, result2.IsNone())

	result3 := LiftA2Option(add, Some(3), None[int]())
	assertTrue(t, result3.IsNone())
}

func TestTransposeOptRes(t *testing.T) {
	// Some(Ok) -> Ok(Some)
	r := TransposeOptRes(Some(Ok(5)))
	assertTrue(t, r.IsOk())
	opt := r.UnwrapOr(None[int]())
	assertEqual(t, opt.UnsafeFromSome(), 5)

	// Some(Err) -> Err
	r2 := TransposeOptRes(Some(Err[int](errors.New("fail"))))
	assertTrue(t, r2.IsErr())

	// None -> Ok(None)
	r3 := TransposeOptRes(None[Result[int]]())
	assertTrue(t, r3.IsOk())
	assertTrue(t, r3.UnwrapOr(Some(0)).IsNone())
}
