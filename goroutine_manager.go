package fn

import "sync"

// GoroutineManager manages the lifecycle of a set of goroutines, ensuring
// they all complete before Stop returns
type GoroutineManager struct {
	wg   sync.WaitGroup
	quit chan struct{}
	mu   sync.Mutex
}

// NewGoroutineManager creates a new GoroutineManager
func NewGoroutineManager() *GoroutineManager {
	return &GoroutineManager{
		quit: make(chan struct{}),
	}
}

// Go launches a goroutine managed by this manager. The provided function
// receives a quit channel it should select on to detect shutdown. Returns
// false if the manager has already been stopped
func (g *GoroutineManager) Go(f func(quit <-chan struct{})) bool {
	g.mu.Lock()
	select {
	case <-g.quit:
		g.mu.Unlock()
		return false
	default:
	}
	g.wg.Add(1)
	g.mu.Unlock()

	go func() {
		defer g.wg.Done()
		f(g.quit)
	}()

	return true
}

// Stop signals all goroutines to stop and waits for them to finish
func (g *GoroutineManager) Stop() {
	g.mu.Lock()
	select {
	case <-g.quit:
		g.mu.Unlock()
		return
	default:
		close(g.quit)
	}
	g.mu.Unlock()

	g.wg.Wait()
}

// Done returns a channel that is closed when Stop is called
func (g *GoroutineManager) Done() <-chan struct{} {
	return g.quit
}
