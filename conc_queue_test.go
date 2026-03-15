package fn

import (
	"testing"
	"time"
)

func TestConcurrentQueueBasic(t *testing.T) {
	q := NewConcurrentQueue[int]()
	q.Start()

	// Send items
	go func() {
		q.ChanIn() <- 1
		q.ChanIn() <- 2
		q.ChanIn() <- 3
	}()

	// Receive in order
	assertEqual(t, <-q.ChanOut(), 1)
	assertEqual(t, <-q.ChanOut(), 2)
	assertEqual(t, <-q.ChanOut(), 3)

	q.Stop()
}

func TestConcurrentQueueBuffering(t *testing.T) {
	q := NewConcurrentQueue[int]()
	q.Start()

	// Send many items quickly (tests overflow buffering)
	done := make(chan struct{})
	go func() {
		for i := 0; i < 100; i++ {
			q.ChanIn() <- i
		}
		close(done)
	}()

	// Wait for all sends
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("sends timed out")
	}

	// Receive all
	for i := 0; i < 100; i++ {
		select {
		case val := <-q.ChanOut():
			assertEqual(t, val, i)
		case <-time.After(time.Second):
			t.Fatalf("timed out at item %d", i)
		}
	}

	q.Stop()
}
