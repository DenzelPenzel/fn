package fn

import (
	"testing"
)

// assertEqual fails the test if got != want
func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

// assertTrue fails the test if the condition is false
func assertTrue(t *testing.T, cond bool) {
	t.Helper()
	if !cond {
		t.Fatal("expected true, got false")
	}
}

// assertFalse fails the test if the condition is true
func assertFalse(t *testing.T, cond bool) {
	t.Helper()
	if cond {
		t.Fatal("expected false, got true")
	}
}

// assertNoError fails the test if err is not nil
func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// assertError fails the test if err is nil
func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

// assertSliceEqual fails if two slices are not element-wise equal
func assertSliceEqual[T comparable](t *testing.T, got, want []T) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("slice length mismatch: got %d, want %d", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("element %d: got %v, want %v", i, got[i], want[i])
		}
	}
}
