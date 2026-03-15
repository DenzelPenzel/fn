package fn

// ConcurrentQueue is an unbounded concurrent queue. Items sent to ChanIn are
// buffered internally (using a List for overflow) and delivered to ChanOut in
// FIFO order
type ConcurrentQueue[T any] struct {
	chanIn  chan T
	chanOut chan T
	quit    chan struct{}
	buf     *List[T]
}

// NewConcurrentQueue creates a new ConcurrentQueue
func NewConcurrentQueue[T any]() *ConcurrentQueue[T] {
	return &ConcurrentQueue[T]{
		chanIn:  make(chan T),
		chanOut: make(chan T),
		quit:    make(chan struct{}),
		buf:     NewList[T](),
	}
}

// ChanIn returns the channel to send items into
func (q *ConcurrentQueue[T]) ChanIn() chan<- T {
	return q.chanIn
}

// ChanOut returns the channel to receive items from
func (q *ConcurrentQueue[T]) ChanOut() <-chan T {
	return q.chanOut
}

// Start begins processing. Must be called before sending/receiving
func (q *ConcurrentQueue[T]) Start() {
	go q.run()
}

// Stop shuts down the queue
func (q *ConcurrentQueue[T]) Stop() {
	close(q.quit)
}

// run is the internal goroutine that shuttles items from chanIn through the
// buffer to chanOut
func (q *ConcurrentQueue[T]) run() {
	for {
		if q.buf.Len() == 0 {
			// Buffer is empty; wait for input or quit
			select {
			case item, ok := <-q.chanIn:
				if !ok {
					return
				}
				q.buf.PushBack(item)

			case <-q.quit:
				return
			}
		} else {
			// Buffer has items; try to send or receive
			front := q.buf.Front().Value
			select {
			case item, ok := <-q.chanIn:
				if !ok {
					return
				}
				q.buf.PushBack(item)

			case q.chanOut <- front:
				q.buf.Remove(q.buf.Front())

			case <-q.quit:
				return
			}
		}
	}
}
