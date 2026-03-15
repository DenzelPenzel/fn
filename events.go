package fn

import "sync"

// EventReceiver is a subscription handle for an EventDistributor
type EventReceiver[T any] struct {
	ch   chan T
	quit chan struct{}
}

// NewEventReceiver creates a new receiver with the given buffer size
func NewEventReceiver[T any](bufSize int) *EventReceiver[T] {
	return &EventReceiver[T]{
		ch:   make(chan T, bufSize),
		quit: make(chan struct{}),
	}
}

// Events returns the channel on which events are delivered
func (r *EventReceiver[T]) Events() <-chan T {
	return r.ch
}

// Stop unsubscribes the receiver. Safe to call multiple times
func (r *EventReceiver[T]) Stop() {
	select {
	case <-r.quit:
	default:
		close(r.quit)
	}
}

// Done returns a channel that is closed when Stop is called
func (r *EventReceiver[T]) Done() <-chan struct{} {
	return r.quit
}

// EventDistributor fans out events to multiple subscribers
type EventDistributor[T any] struct {
	mu          sync.RWMutex
	subscribers map[*EventReceiver[T]]struct{}
}

// NewEventDistributor creates a new distributor
func NewEventDistributor[T any]() *EventDistributor[T] {
	return &EventDistributor[T]{
		subscribers: make(map[*EventReceiver[T]]struct{}),
	}
}

// Subscribe adds a receiver and returns it
func (d *EventDistributor[T]) Subscribe(
	r *EventReceiver[T],
) *EventReceiver[T] {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.subscribers[r] = struct{}{}
	return r
}

// Unsubscribe removes a receiver
func (d *EventDistributor[T]) Unsubscribe(r *EventReceiver[T]) {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.subscribers, r)
}

// NotifySubscribers sends the event to all current subscribers. Subscribers
// whose channels are full or who have stopped are skipped
func (d *EventDistributor[T]) NotifySubscribers(event T) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	for sub := range d.subscribers {
		select {
		case <-sub.quit:
			continue
		default:
		}

		select {
		case sub.ch <- event:
		default:
			// Drop event if subscriber buffer is full
		}
	}
}

// NumSubscribers returns the current subscriber count
func (d *EventDistributor[T]) NumSubscribers() int {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return len(d.subscribers)
}
