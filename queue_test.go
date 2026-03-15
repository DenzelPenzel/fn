package fn

import "testing"

func TestQueueBasic(t *testing.T) {
	q := NewQueue[int]()
	assertTrue(t, q.IsEmpty())
	assertEqual(t, q.Size(), 0)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assertEqual(t, q.Size(), 3)
	assertFalse(t, q.IsEmpty())

	// Peek
	assertEqual(t, q.Peek().UnsafeFromSome(), 1)

	// Dequeue in FIFO order
	assertEqual(t, q.Dequeue().UnsafeFromSome(), 1)
	assertEqual(t, q.Dequeue().UnsafeFromSome(), 2)
	assertEqual(t, q.Dequeue().UnsafeFromSome(), 3)
	assertTrue(t, q.Dequeue().IsNone())
	assertTrue(t, q.IsEmpty())
}

func TestQueueEmpty(t *testing.T) {
	q := NewQueue[string]()
	assertTrue(t, q.Dequeue().IsNone())
	assertTrue(t, q.Peek().IsNone())
}
