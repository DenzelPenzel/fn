package fn

import (
	"testing"
	"time"
)

func TestGuardTestPasses(t *testing.T) {
	GuardTest(t, time.Second, func() {
		// Fast test that completes in time
	})
}

func TestGuardTestCallsFunction(t *testing.T) {
	var called bool
	GuardTest(t, time.Second, func() {
		called = true
	})
	assertTrue(t, called)
}
