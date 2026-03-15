package fn

import "time"

// RecvOrTimeout receives from ch or returns None after the timeout
func RecvOrTimeout[T any](ch <-chan T, timeout time.Duration) Option[T] {
	select {
	case val := <-ch:
		return Some(val)
	case <-time.After(timeout):
		return None[T]()
	}
}

// RecvResp receives from ch or returns None if quit is closed
func RecvResp[T any](ch <-chan T, quit <-chan struct{}) Option[T] {
	select {
	case val := <-ch:
		return Some(val)
	case <-quit:
		return None[T]()
	}
}
