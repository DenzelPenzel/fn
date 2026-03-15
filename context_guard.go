package fn

import (
	"context"
	"sync"
)

// ContextGuardOption configures a ContextGuard
type ContextGuardOption func(*ContextGuard)

// WithCustomContext sets a parent context for the guard
func WithCustomContext(ctx context.Context) ContextGuardOption {
	return func(g *ContextGuard) {
		g.ctx, g.cancel = context.WithCancel(ctx)
	}
}

// ContextGuard manages a context-based lifecycle, ensuring spawned goroutines
// respect the parent context and can be cleanly shut down
type ContextGuard struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
	mu     sync.Mutex
}

// NewContextGuard creates a new ContextGuard with optional configuration
func NewContextGuard(opts ...ContextGuardOption) *ContextGuard {
	g := &ContextGuard{}
	for _, opt := range opts {
		opt(g)
	}

	// Default to background context if no custom context was provided
	if g.ctx == nil {
		g.ctx, g.cancel = context.WithCancel(context.Background())
	}

	return g
}

// Create returns a derived context and a WaitGroup done function. The caller
// should defer the done function to signal completion
func (g *ContextGuard) Create(
	ctx context.Context,
) (context.Context, func()) {
	g.mu.Lock()
	defer g.mu.Unlock()

	derived, cancel := context.WithCancel(ctx)
	g.wg.Add(1)

	done := func() {
		cancel()
		g.wg.Done()
	}

	return derived, done
}

// Quit signals all derived contexts to cancel and waits for completion
func (g *ContextGuard) Quit() {
	g.cancel()
	g.wg.Wait()
}

// Done returns a channel closed when the guard's context is cancelled
func (g *ContextGuard) Done() <-chan struct{} {
	return g.ctx.Done()
}

// Context returns the guard's context
func (g *ContextGuard) Context() context.Context {
	return g.ctx
}
