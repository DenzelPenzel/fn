package fn

// SendOrQuit sends val on ch or returns false if quit is closed first
func SendOrQuit[T any](ch chan<- T, val T, quit <-chan struct{}) bool {
	select {
	case ch <- val:
		return true
	case <-quit:
		return false
	}
}
