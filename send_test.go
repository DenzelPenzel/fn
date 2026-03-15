package fn

import "testing"

func TestSendOrQuit(t *testing.T) {
	ch := make(chan int, 1)
	quit := make(chan struct{})

	ok := SendOrQuit(ch, 42, quit)
	assertTrue(t, ok)
	assertEqual(t, <-ch, 42)

	// Quit before send
	ch2 := make(chan int) // Unbuffered, will block.
	close(quit)
	ok2 := SendOrQuit(ch2, 1, quit)
	assertFalse(t, ok2)
}
