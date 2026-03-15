package fn

import (
	"errors"
	"fmt"
	"testing"
)

func TestComp(t *testing.T) {
	double := func(x int) int { return x * 2 }
	addOne := func(x int) int { return x + 1 }

	// Comp(double, addOne)(3) == double(addOne(3)) == 8
	composed := Comp(double, addOne)
	assertEqual(t, composed(3), 8)
}

func TestIden(t *testing.T) {
	assertEqual(t, Iden(42), 42)
	assertEqual(t, Iden("hello"), "hello")
}

func TestConst(t *testing.T) {
	always5 := Const[int, string](5)
	assertEqual(t, always5("anything"), 5)
	assertEqual(t, always5(""), 5)
}

func TestEqNeq(t *testing.T) {
	assertTrue(t, Eq(1, 1))
	assertFalse(t, Eq(1, 2))
	assertFalse(t, Neq(1, 1))
	assertTrue(t, Neq(1, 2))
}

type copyInt int

func (c copyInt) Copy() copyInt {
	return c
}

func TestCopyAll(t *testing.T) {
	xs := []copyInt{1, 2, 3}
	result := CopyAll[copyInt](xs)
	assertSliceEqual(t, result, xs)
}

type copyErrInt struct {
	val int
	err error
}

func (c copyErrInt) Copy() (copyErrInt, error) {
	return copyErrInt{val: c.val}, c.err
}

func TestCopyAllErr(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		xs := []copyErrInt{{val: 1}, {val: 2}}
		result, err := CopyAllErr[copyErrInt](xs)
		assertNoError(t, err)
		assertEqual(t, len(result), 2)
	})

	t.Run("error", func(t *testing.T) {
		xs := []copyErrInt{
			{val: 1},
			{val: 2, err: fmt.Errorf("fail")},
		}
		_, err := CopyAllErr[copyErrInt](xs)
		assertError(t, err)
	})
}

// Verify that errors.Is works with our test error
func TestCopyAllErrPropagation(t *testing.T) {
	sentinel := errors.New("sentinel")
	xs := []copyErrInt{{val: 1, err: sentinel}}
	_, err := CopyAllErr[copyErrInt](xs)
	assertTrue(t, errors.Is(err, sentinel))
}
