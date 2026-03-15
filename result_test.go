package fn

import (
	"errors"
	"fmt"
	"testing"
)

func TestResultOk(t *testing.T) {
	r := Ok(42)
	assertTrue(t, r.IsOk())
	assertFalse(t, r.IsErr())
	val, err := r.Unpack()
	assertNoError(t, err)
	assertEqual(t, val, 42)
}

func TestResultErr(t *testing.T) {
	r := Err[int](errors.New("boom"))
	assertFalse(t, r.IsOk())
	assertTrue(t, r.IsErr())
	_, err := r.Unpack()
	assertError(t, err)
}

func TestResultErrf(t *testing.T) {
	r := Errf[int]("code %d", 404)
	assertTrue(t, r.IsErr())
	assertEqual(t, r.Err().Error(), "code 404")
}

func TestNewResult(t *testing.T) {
	r := NewResult(42, nil)
	assertTrue(t, r.IsOk())
	assertEqual(t, r.UnwrapOr(0), 42)

	r2 := NewResult(0, errors.New("err"))
	assertTrue(t, r2.IsErr())
}

func TestResultUnpack(t *testing.T) {
	r := Ok("hello")
	val, err := r.Unpack()
	assertNoError(t, err)
	assertEqual(t, val, "hello")

	r2 := Err[string](errors.New("fail"))
	val2, err2 := r2.Unpack()
	assertError(t, err2)
	assertEqual(t, val2, "")
}

func TestResultMapOk(t *testing.T) {
	r := Ok(5).MapOk(func(i int) int { return i * 2 })
	assertEqual(t, r.UnwrapOr(0), 10)

	r2 := Err[int](errors.New("e")).MapOk(func(i int) int { return i * 2 })
	assertTrue(t, r2.IsErr())
}

func TestResultMapErr(t *testing.T) {
	r := Err[int](errors.New("e")).MapErr(func(e error) error {
		return fmt.Errorf("wrapped: %w", e)
	})
	assertTrue(t, r.IsErr())

	r2 := Ok(5).MapErr(func(e error) error { return e })
	assertTrue(t, r2.IsOk())
}

func TestResultOkToSome(t *testing.T) {
	assertTrue(t, Ok(5).OkToSome().IsSome())
	assertTrue(t, Err[int](errors.New("e")).OkToSome().IsNone())
}

func TestResultWhenOkErr(t *testing.T) {
	var okCalled, errCalled bool
	Ok(1).WhenOk(func(_ int) { okCalled = true })
	assertTrue(t, okCalled)

	Err[int](errors.New("e")).WhenErr(func(_ error) { errCalled = true })
	assertTrue(t, errCalled)
}

func TestResultUnwrapOr(t *testing.T) {
	assertEqual(t, Ok(5).UnwrapOr(0), 5)
	assertEqual(t, Err[int](errors.New("e")).UnwrapOr(99), 99)
}

func TestResultUnwrapOrElse(t *testing.T) {
	assertEqual(t, Ok(5).UnwrapOrElse(func() int { return 0 }), 5)
	assertEqual(t, Err[int](errors.New("e")).UnwrapOrElse(func() int { return 99 }), 99)
}

func TestResultFlatMap(t *testing.T) {
	r := Ok(5).FlatMap(func(i int) Result[int] {
		if i > 0 {
			return Ok(i * 10)
		}
		return Errf[int]("negative")
	})
	assertEqual(t, r.UnwrapOr(0), 50)

	r2 := Err[int](errors.New("e")).FlatMap(func(_ int) Result[int] {
		return Ok(0)
	})
	assertTrue(t, r2.IsErr())
}

func TestResultOrElse(t *testing.T) {
	r := Err[int](errors.New("e")).OrElse(func(_ error) Result[int] {
		return Ok(42)
	})
	assertEqual(t, r.UnwrapOr(0), 42)

	r2 := Ok(5).OrElse(func(_ error) Result[int] { return Ok(0) })
	assertEqual(t, r2.UnwrapOr(0), 5)
}

func TestResultSink(t *testing.T) {
	var gotOk int
	var gotErr error
	Ok(5).Sink(func(v int) { gotOk = v }, func(e error) { gotErr = e })
	assertEqual(t, gotOk, 5)

	Err[int](errors.New("boom")).Sink(
		func(_ int) {},
		func(e error) { gotErr = e },
	)
	assertError(t, gotErr)
}

func TestMapResultOk(t *testing.T) {
	r := MapResultOk(func(_ int) string { return "ok" }, Ok(5))
	assertTrue(t, r.IsOk())
	assertEqual(t, r.UnwrapOr(""), "ok")

	r2 := MapResultOk(func(_ int) string { return "ok" }, Err[int](errors.New("e")))
	assertTrue(t, r2.IsErr())
}

func TestFlatMapResult(t *testing.T) {
	r := FlatMapResult(func(i int) Result[string] {
		return Ok(fmt.Sprintf("%d", i))
	}, Ok(42))
	assertEqual(t, r.UnwrapOr(""), "42")
}

func TestFlattenResult(t *testing.T) {
	r := FlattenResult(Ok(Ok(5)))
	assertEqual(t, r.UnwrapOr(0), 5)

	r2 := FlattenResult(Ok(Err[int](errors.New("inner"))))
	assertTrue(t, r2.IsErr())

	r3 := FlattenResult(Err[Result[int]](errors.New("outer")))
	assertTrue(t, r3.IsErr())
}

func TestLiftA2Result(t *testing.T) {
	add := func(a, b int) int { return a + b }
	r := LiftA2Result(add, Ok(3), Ok(4))
	assertEqual(t, r.UnwrapOr(0), 7)

	r2 := LiftA2Result(add, Err[int](errors.New("a")), Ok(4))
	assertTrue(t, r2.IsErr())

	r3 := LiftA2Result(add, Ok(3), Err[int](errors.New("b")))
	assertTrue(t, r3.IsErr())
}

func TestTransposeResOpt(t *testing.T) {
	// Ok(Some) -> Some(Ok)
	opt := TransposeResOpt(Ok(Some(5)))
	assertTrue(t, opt.IsSome())
	assertTrue(t, opt.UnsafeFromSome().IsOk())

	// Ok(None) -> None
	opt2 := TransposeResOpt(Ok(None[int]()))
	assertTrue(t, opt2.IsNone())

	// Err -> Some(Err)
	opt3 := TransposeResOpt(Err[Option[int]](errors.New("e")))
	assertTrue(t, opt3.IsSome())
	assertTrue(t, opt3.UnsafeFromSome().IsErr())
}
