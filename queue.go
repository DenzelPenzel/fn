package fn

// Queue is a simple FIFO queue backed by a slice
type Queue[T any] struct {
	items []T
}

// NewQueue creates a new empty queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue adds an element to the back of the queue
func (q *Queue[T]) Enqueue(val T) {
	q.items = append(q.items, val)
}

// Dequeue removes and returns the front element
func (q *Queue[T]) Dequeue() Option[T] {
	if len(q.items) == 0 {
		return None[T]()
	}
	val := q.items[0]
	q.items = q.items[1:]
	return Some(val)
}

// Peek returns the front element without removing it
func (q *Queue[T]) Peek() Option[T] {
	if len(q.items) == 0 {
		return None[T]()
	}
	return Some(q.items[0])
}

// IsEmpty returns true if the queue has no elements
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of elements in the queue
func (q *Queue[T]) Size() int {
	return len(q.items)
}
