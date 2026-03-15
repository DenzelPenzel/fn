package fn_test

import (
	"errors"
	"fmt"
	"strings"

	"github.com/denzelpenzel/fn"
)

func ExampleSome() {
	o := fn.Some(42)
	fmt.Println(o.IsSome(), o.UnwrapOr(0))
	// Output: true 42
}

func ExampleNone() {
	o := fn.None[int]()
	fmt.Println(o.IsNone(), o.UnwrapOr(-1))
	// Output: true -1
}

func ExampleMapOption() {
	o := fn.Some(3)
	doubled := fn.MapOption(func(x int) int { return x * 2 }, o)
	fmt.Println(doubled.UnwrapOr(0))
	// Output: 6
}

func ExampleFlatMapOption() {
	safeSqrt := func(x int) fn.Option[int] {
		if x < 0 {
			return fn.None[int]()
		}
		return fn.Some(x * x)
	}
	result := fn.FlatMapOption(safeSqrt, fn.Some(5))
	fmt.Println(result.UnwrapOr(0))
	// Output: 25
}

func ExampleOk() {
	r := fn.Ok(100)
	fmt.Println(r.IsOk(), r.UnwrapOr(0))
	// Output: true 100
}

func ExampleErr() {
	r := fn.Err[int](errors.New("fail"))
	fmt.Println(r.IsErr(), r.UnwrapOr(-1))
	// Output: true -1
}

func ExampleNewResult() {
	r := fn.NewResult(42, nil)
	fmt.Println(r.IsOk(), r.UnwrapOr(0))
	// Output: true 42
}

func ExampleResult_FlatMap() {
	double := func(x int) fn.Result[int] {
		if x > 100 {
			return fn.Err[int](errors.New("too large"))
		}
		return fn.Ok(x * 2)
	}
	r := fn.Ok(10).FlatMap(double)
	fmt.Println(r.UnwrapOr(0))
	// Output: 20
}

func ExampleMap() {
	result := fn.Map(strings.ToUpper, []string{"hello", "world"})
	fmt.Println(result)
	// Output: [HELLO WORLD]
}

func ExampleFilter() {
	evens := fn.Filter(func(x int) bool {
		return x%2 == 0
	}, []int{1, 2, 3, 4, 5, 6})
	fmt.Println(evens)
	// Output: [2 4 6]
}

func ExampleFoldl() {
	sum := fn.Foldl(func(acc, x int) int {
		return acc + x
	}, 0, []int{1, 2, 3, 4})
	fmt.Println(sum)
	// Output: 10
}

func ExampleFind() {
	result := fn.Find(func(x int) bool {
		return x > 3
	}, []int{1, 2, 3, 4, 5})
	fmt.Println(result.UnwrapOr(0))
	// Output: 4
}

func ExampleNewSet() {
	s := fn.NewSet(1, 2, 3, 2, 1)
	fmt.Println(s.Size(), s.Contains(2), s.Contains(4))
	// Output: 3 true false
}

func ExampleComp() {
	double := func(x int) int { return x * 2 }
	addOne := func(x int) int { return x + 1 }
	doubleThenAdd := fn.Comp(addOne, double)
	fmt.Println(doubleThenAdd(5))
	// Output: 11
}

func ExampleCollectResults() {
	rs := []fn.Result[int]{fn.Ok(1), fn.Ok(2), fn.Ok(3)}
	all := fn.CollectResults(rs)
	val, err := all.Unpack()
	fmt.Println(val, err)
	// Output: [1 2 3] <nil>
}
