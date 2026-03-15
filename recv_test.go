package fn

import (
	"testing"
	"time"
)

func TestRecvOrTimeout(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 42
	result := RecvOrTimeout(ch, time.Second)
	assertEqual(t, result.UnsafeFromSome(), 42)

	// Timeout case
	ch2 := make(chan int)
	result2 := RecvOrTimeout(ch2, 10*time.Millisecond)
	assertTrue(t, result2.IsNone())
}

func TestRecvResp(t *testing.T) {
	ch := make(chan int, 1)
	quit := make(chan struct{})
	ch <- 10
	result := RecvResp(ch, quit)
	assertEqual(t, result.UnsafeFromSome(), 10)

	// Quit case
	ch2 := make(chan int)
	close(quit)
	result2 := RecvResp(ch2, quit)
	assertTrue(t, result2.IsNone())
}
