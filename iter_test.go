package fn

import (
	"errors"
	"iter"
	"testing"
)

func TestCollect(t *testing.T) {
	seq := func(yield func(int) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i) {
				return
			}
		}
	}
	result := Collect(iter.Seq[int](seq))
	assertSliceEqual(t, result, []int{1, 2, 3})
}

func TestCollectErr(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		seq := func(yield func(int, error) bool) {
			for i := 1; i <= 3; i++ {
				if !yield(i, nil) {
					return
				}
			}
		}
		result, err := CollectErr(iter.Seq2[int, error](seq))
		assertNoError(t, err)
		assertSliceEqual(t, result, []int{1, 2, 3})
	})

	t.Run("error", func(t *testing.T) {
		sentinel := errors.New("stop")
		seq := func(yield func(int, error) bool) {
			if !yield(1, nil) {
				return
			}
			if !yield(0, sentinel) {
				return
			}
			yield(3, nil) // Should not be reached.
		}
		_, err := CollectErr(iter.Seq2[int, error](seq))
		assertTrue(t, errors.Is(err, sentinel))
	})
}
