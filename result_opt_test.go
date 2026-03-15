package fn

import (
	"errors"
	"testing"
)

func TestResultOptOkOpt(t *testing.T) {
	ro := OkOpt(42)
	assertTrue(t, ro.IsSome())
	assertFalse(t, ro.IsNone())
	assertFalse(t, ro.IsErr())
}

func TestResultOptNoneOpt(t *testing.T) {
	ro := NoneOpt[int]()
	assertFalse(t, ro.IsSome())
	assertTrue(t, ro.IsNone())
	assertFalse(t, ro.IsErr())
}

func TestResultOptErrOpt(t *testing.T) {
	ro := ErrOpt[int](errors.New("boom"))
	assertFalse(t, ro.IsSome())
	assertFalse(t, ro.IsNone())
	assertTrue(t, ro.IsErr())
}

func TestMapResultOpt(t *testing.T) {
	ro := OkOpt(5)
	mapped := MapResultOpt(func(_ int) string { return "ok" }, ro)
	assertTrue(t, mapped.IsSome())

	ro2 := NoneOpt[int]()
	mapped2 := MapResultOpt(func(_ int) string { return "ok" }, ro2)
	assertTrue(t, mapped2.IsNone())

	ro3 := ErrOpt[int](errors.New("e"))
	mapped3 := MapResultOpt(func(_ int) string { return "ok" }, ro3)
	assertTrue(t, mapped3.IsErr())
}

func TestAndThenResultOpt(t *testing.T) {
	ro := OkOpt(5)
	result := AndThenResultOpt(ro, func(_ int) ResultOpt[string] {
		return OkOpt("five")
	})
	assertTrue(t, result.IsSome())

	ro2 := NoneOpt[int]()
	result2 := AndThenResultOpt(ro2, func(_ int) ResultOpt[string] {
		return OkOpt("x")
	})
	assertTrue(t, result2.IsNone())

	ro3 := ErrOpt[int](errors.New("e"))
	result3 := AndThenResultOpt(ro3, func(_ int) ResultOpt[string] {
		return OkOpt("x")
	})
	assertTrue(t, result3.IsErr())
}
