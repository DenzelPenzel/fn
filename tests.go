package fn

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// GuardTest wraps a test with a timeout. If the test does not complete within
// the given duration, it dumps all goroutine stacks and fails
func GuardTest(t *testing.T, timeout time.Duration, f func()) {
	t.Helper()

	done := make(chan struct{})
	go func() {
		defer close(done)
		f()
	}()

	select {
	case <-done:
	case <-time.After(timeout):
		buf := make([]byte, 1<<20)
		n := runtime.Stack(buf, true)
		t.Fatalf(
			"test timed out after %v\n\ngoroutine dump:\n%s",
			timeout, fmt.Sprintf("%s", buf[:n]),
		)
	}
}
