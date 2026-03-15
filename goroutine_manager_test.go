package fn

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestGoroutineManagerBasic(t *testing.T) {
	gm := NewGoroutineManager()
	var count atomic.Int32

	for i := 0; i < 5; i++ {
		ok := gm.Go(func(quit <-chan struct{}) {
			count.Add(1)
			<-quit
		})
		assertTrue(t, ok)
	}

	// Give goroutines time to start
	time.Sleep(10 * time.Millisecond)
	assertEqual(t, count.Load(), int32(5))

	gm.Stop()

	// After stop, Go should return false
	ok := gm.Go(func(quit <-chan struct{}) {})
	assertFalse(t, ok)
}

func TestGoroutineManagerDone(t *testing.T) {
	gm := NewGoroutineManager()
	select {
	case <-gm.Done():
		t.Fatal("Done should not be closed yet")
	default:
	}

	gm.Stop()

	select {
	case <-gm.Done():
	case <-time.After(time.Second):
		t.Fatal("Done should be closed after Stop")
	}
}

func TestGoroutineManagerDoubleStop(t *testing.T) {
	gm := NewGoroutineManager()
	gm.Stop()
	gm.Stop() // Should not panic.
}
