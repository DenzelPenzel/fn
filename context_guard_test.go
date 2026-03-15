package fn

import (
	"context"
	"testing"
	"time"
)

func TestContextGuardBasic(t *testing.T) {
	g := NewContextGuard()

	ctx, done := g.Create(g.Context())

	// Context should be alive
	select {
	case <-ctx.Done():
		t.Fatal("context should not be done")
	default:
	}

	// Quit cancels all derived contexts, but we need to release our hold
	// first from a goroutine so Quit can complete
	go func() {
		<-ctx.Done()
		done()
	}()

	g.Quit()

	select {
	case <-g.Done():
	default:
		t.Fatal("guard Done should be closed after Quit")
	}
}

func TestContextGuardWithCustomContext(t *testing.T) {
	parent, cancel := context.WithCancel(context.Background())

	g := NewContextGuard(WithCustomContext(parent))
	ctx, done := g.Create(g.Context())

	// Cancelling parent should propagate
	cancel()

	select {
	case <-ctx.Done():
	case <-time.After(time.Second):
		t.Fatal("context should be cancelled")
	}

	// Release hold so Quit can return
	done()
	g.Quit()
}

func TestContextGuardMultiple(t *testing.T) {
	g := NewContextGuard()
	var dones []func()

	for range 3 {
		ctx, done := g.Create(g.Context())
		dones = append(dones, done)

		// Release hold when context is cancelled
		go func() {
			<-ctx.Done()
		}()
	}

	// Release all holds then quit
	for _, done := range dones {
		done()
	}
	g.Quit()

	select {
	case <-g.Done():
	default:
		t.Fatal("guard Done should be closed after Quit")
	}
}
