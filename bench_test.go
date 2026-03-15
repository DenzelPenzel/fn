package fn_test

import (
	"testing"

	"github.com/denzelpenzel/fn"
)

func benchSlice(size int) []int {
	xs := make([]int, size)
	for i := range xs {
		xs[i] = i
	}
	return xs
}

func BenchmarkMap(b *testing.B) {
	for _, size := range []int{10, 100, 1000} {
		xs := benchSlice(size)
		double := func(x int) int { return x * 2 }
		b.Run(
			sizeLabel(size), func(b *testing.B) {
				b.ReportAllocs()
				for b.Loop() {
					fn.Map(double, xs)
				}
			},
		)
	}
}

func BenchmarkFilter(b *testing.B) {
	for _, size := range []int{10, 100, 1000} {
		xs := benchSlice(size)
		even := func(x int) bool { return x%2 == 0 }
		b.Run(
			sizeLabel(size), func(b *testing.B) {
				b.ReportAllocs()
				for b.Loop() {
					fn.Filter(even, xs)
				}
			},
		)
	}
}

func BenchmarkFoldl(b *testing.B) {
	xs := benchSlice(1000)
	add := func(a, x int) int { return a + x }
	b.ReportAllocs()
	for b.Loop() {
		fn.Foldl(add, 0, xs)
	}
}

func BenchmarkFind(b *testing.B) {
	xs := benchSlice(1000)
	// Target near the end to measure worst case
	pred := func(x int) bool { return x == 999 }
	b.ReportAllocs()
	for b.Loop() {
		fn.Find(pred, xs)
	}
}

func BenchmarkMapOption(b *testing.B) {
	o := fn.Some(42)
	double := func(x int) int { return x * 2 }
	b.ReportAllocs()
	for b.Loop() {
		fn.MapOption(double, o)
	}
}

func BenchmarkResultMapOk(b *testing.B) {
	r := fn.Ok(42)
	double := func(x int) int { return x * 2 }
	b.ReportAllocs()
	for b.Loop() {
		r.MapOk(double)
	}
}

func BenchmarkSetContains(b *testing.B) {
	s := fn.NewSet(benchSlice(1000)...)
	b.ReportAllocs()
	for b.Loop() {
		s.Contains(500)
	}
}

func BenchmarkSetUnion(b *testing.B) {
	a := fn.NewSet(benchSlice(500)...)
	bSet := fn.NewSet(benchSlice(500)...)
	b.ReportAllocs()
	for b.Loop() {
		a.Union(bSet)
	}
}

func BenchmarkConcurrentQueue(b *testing.B) {
	q := fn.NewConcurrentQueue[int]()
	q.Start()
	defer q.Stop()

	b.ReportAllocs()
	for b.Loop() {
		q.ChanIn() <- 1
		<-q.ChanOut()
	}
}

func sizeLabel(n int) string {
	switch {
	case n >= 1000:
		return "1000"
	case n >= 100:
		return "100"
	default:
		return "10"
	}
}
