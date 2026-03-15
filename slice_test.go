package fn

import (
	"errors"
	"sync/atomic"
	"testing"
)

func TestMap(t *testing.T) {
	result := Map(func(i int) int { return i * 2 }, []int{1, 2, 3})
	assertSliceEqual(t, result, []int{2, 4, 6})

	// Empty slice
	assertSliceEqual(t, Map(func(i int) int { return i }, []int(nil)), []int{})
}

func TestFilter(t *testing.T) {
	even := func(i int) bool { return i%2 == 0 }
	result := Filter(even, []int{1, 2, 3, 4, 5})
	assertSliceEqual(t, result, []int{2, 4})

	// None match
	result2 := Filter(even, []int{1, 3, 5})
	assertEqual(t, len(result2), 0)
}

func TestFilterMap(t *testing.T) {
	f := func(i int) Option[string] {
		if i > 2 {
			return Some("big")
		}
		return None[string]()
	}
	result := FilterMap(f, []int{1, 2, 3, 4})
	assertSliceEqual(t, result, []string{"big", "big"})
}

func TestTrimNones(t *testing.T) {
	xs := []Option[int]{Some(1), None[int](), Some(3)}
	result := TrimNones(xs)
	assertSliceEqual(t, result, []int{1, 3})
}

func TestFoldl(t *testing.T) {
	// Left fold: ((0-1)-2)-3 = -6
	result := Foldl(func(acc, x int) int { return acc - x }, 0, []int{1, 2, 3})
	assertEqual(t, result, -6)
}

func TestFoldr(t *testing.T) {
	// Right fold with string concat: 1:(2:(3:end)) = "123end"
	result := Foldr(
		func(x int, acc string) string {
			return string(rune('0'+x)) + acc
		}, "end", []int{1, 2, 3},
	)
	assertEqual(t, result, "123end")
}

func TestSum(t *testing.T) {
	assertEqual(t, Sum([]int{1, 2, 3}), 6)
	assertEqual(t, Sum([]float64{1.5, 2.5}), 4.0)
	assertEqual(t, Sum([]int(nil)), 0)
}

func TestFind(t *testing.T) {
	result := Find(func(i int) bool { return i > 3 }, []int{1, 2, 4, 5})
	assertEqual(t, result.UnsafeFromSome(), 4)

	result2 := Find(func(i int) bool { return i > 10 }, []int{1, 2, 3})
	assertTrue(t, result2.IsNone())
}

func TestFindIdx(t *testing.T) {
	result := FindIdx(func(i int) bool { return i > 3 }, []int{1, 2, 4, 5})
	idx, val := result.UnsafeFromSome().Unpack()
	assertEqual(t, idx, 2)
	assertEqual(t, val, 4)
}

func TestElem(t *testing.T) {
	assertTrue(t, Elem(3, []int{1, 2, 3}))
	assertFalse(t, Elem(4, []int{1, 2, 3}))
}

func TestAllAny(t *testing.T) {
	pos := func(i int) bool { return i > 0 }
	assertTrue(t, All(pos, []int{1, 2, 3}))
	assertFalse(t, All(pos, []int{1, -1, 3}))
	assertTrue(t, All(pos, []int(nil)))

	assertTrue(t, Any(pos, []int{-1, 0, 1}))
	assertFalse(t, Any(pos, []int{-1, -2}))
	assertFalse(t, Any(pos, []int(nil)))
}

func TestFlatten(t *testing.T) {
	result := Flatten([][]int{{1, 2}, {3}, {4, 5}})
	assertSliceEqual(t, result, []int{1, 2, 3, 4, 5})
}

func TestReplicate(t *testing.T) {
	result := Replicate(3, "x")
	assertSliceEqual(t, result, []string{"x", "x", "x"})
	assertSliceEqual(t, Replicate(0, 1), []int{})
}

func TestSpan(t *testing.T) {
	pre, post := Span(func(i int) bool { return i < 3 }, []int{1, 2, 3, 4})
	assertSliceEqual(t, pre, []int{1, 2})
	assertSliceEqual(t, post, []int{3, 4})

	pre2, post2 := Span(func(i int) bool { return i < 10 }, []int{1, 2, 3})
	assertSliceEqual(t, pre2, []int{1, 2, 3})
	assertEqual(t, len(post2), 0)
}

func TestSplitAt(t *testing.T) {
	a, b := SplitAt(2, []int{1, 2, 3, 4})
	assertSliceEqual(t, a, []int{1, 2})
	assertSliceEqual(t, b, []int{3, 4})

	// Out of bounds
	a2, b2 := SplitAt(10, []int{1, 2})
	assertSliceEqual(t, a2, []int{1, 2})
	assertEqual(t, len(b2), 0)

	a3, b3 := SplitAt(-1, []int{1, 2})
	assertEqual(t, len(a3), 0)
	assertSliceEqual(t, b3, []int{1, 2})
}

func TestZipWith(t *testing.T) {
	result := ZipWith(
		func(a, b int) int { return a + b },
		[]int{1, 2, 3}, []int{10, 20},
	)
	assertSliceEqual(t, result, []int{11, 22})
}

func TestHeadTailInitLast(t *testing.T) {
	xs := []int{1, 2, 3}
	assertEqual(t, Head(xs).UnsafeFromSome(), 1)
	assertSliceEqual(t, Tail(xs).UnsafeFromSome(), []int{2, 3})
	assertSliceEqual(t, Init(xs).UnsafeFromSome(), []int{1, 2})
	assertEqual(t, Last(xs).UnsafeFromSome(), 3)

	// Empty
	assertTrue(t, Head([]int(nil)).IsNone())
	assertTrue(t, Tail([]int(nil)).IsNone())
	assertTrue(t, Init([]int(nil)).IsNone())
	assertTrue(t, Last([]int(nil)).IsNone())
}

func TestUnconsUnsnoc(t *testing.T) {
	xs := []int{1, 2, 3}
	uc := Uncons(xs).UnsafeFromSome()
	assertEqual(t, uc.First(), 1)
	assertSliceEqual(t, uc.Second(), []int{2, 3})

	us := Unsnoc(xs).UnsafeFromSome()
	assertSliceEqual(t, us.First(), []int{1, 2})
	assertEqual(t, us.Second(), 3)

	assertTrue(t, Uncons([]int(nil)).IsNone())
	assertTrue(t, Unsnoc([]int(nil)).IsNone())
}

func TestSliceToMap(t *testing.T) {
	xs := []T2[string, int]{NewT2("a", 1), NewT2("b", 2)}
	m := SliceToMap(xs)
	assertEqual(t, m["a"], 1)
	assertEqual(t, m["b"], 2)
}

func TestHasDuplicates(t *testing.T) {
	assertTrue(t, HasDuplicates([]int{1, 2, 1}))
	assertFalse(t, HasDuplicates([]int{1, 2, 3}))
	assertFalse(t, HasDuplicates([]int(nil)))
}

func TestLen(t *testing.T) {
	assertEqual(t, Len([]int{1, 2, 3}), 3)
	assertEqual(t, Len([]int(nil)), 0)
}

func TestForEachConc(t *testing.T) {
	var counter atomic.Int64
	ForEachConc(func(i int) {
		counter.Add(int64(i))
	}, []int{1, 2, 3, 4})
	assertEqual(t, counter.Load(), int64(10))
}

func TestCollectOptions(t *testing.T) {
	opts := []Option[int]{Some(1), Some(2), Some(3)}
	result := CollectOptions(opts)
	assertTrue(t, result.IsSome())
	assertSliceEqual(t, result.UnsafeFromSome(), []int{1, 2, 3})

	opts2 := []Option[int]{Some(1), None[int](), Some(3)}
	assertTrue(t, CollectOptions(opts2).IsNone())

	// Empty slice
	assertTrue(t, CollectOptions([]Option[int]{}).IsSome())
}

func TestCollectResults(t *testing.T) {
	rs := []Result[int]{Ok(1), Ok(2), Ok(3)}
	result := CollectResults(rs)
	assertTrue(t, result.IsOk())
	assertSliceEqual(t, result.UnwrapOr(nil), []int{1, 2, 3})

	rs2 := []Result[int]{Ok(1), Err[int](errors.New("e")), Ok(3)}
	assertTrue(t, CollectResults(rs2).IsErr())
}

func TestTraverseOption(t *testing.T) {
	f := func(i int) Option[int] {
		if i > 0 {
			return Some(i * 10)
		}
		return None[int]()
	}

	result := TraverseOption(f, []int{1, 2, 3})
	assertTrue(t, result.IsSome())
	assertSliceEqual(t, result.UnsafeFromSome(), []int{10, 20, 30})

	result2 := TraverseOption(f, []int{1, -1, 3})
	assertTrue(t, result2.IsNone())
}

func TestTraverseResult(t *testing.T) {
	f := func(i int) Result[int] {
		if i > 0 {
			return Ok(i * 10)
		}
		return Errf[int]("negative: %d", i)
	}

	result := TraverseResult(f, []int{1, 2, 3})
	assertTrue(t, result.IsOk())
	assertSliceEqual(t, result.UnwrapOr(nil), []int{10, 20, 30})

	result2 := TraverseResult(f, []int{1, -1, 3})
	assertTrue(t, result2.IsErr())
}
