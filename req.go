package fn

// Req is a request/response pair sent over a channel. The caller sends an
// Input and waits for an Output on the response channel
type Req[Input, Output any] struct {
	Input Input
	resp  chan Output
}

// NewReq creates a Req with the given input
func NewReq[Input, Output any](input Input) Req[Input, Output] {
	return Req[Input, Output]{
		Input: input,
		resp:  make(chan Output, 1),
	}
}

// Dispatch sends a request on ch and blocks until a response arrives or quit
// is closed
func Dispatch[Input, Output any](
	req Req[Input, Output], ch chan<- Req[Input, Output],
	quit <-chan struct{},
) Option[Output] {
	select {
	case ch <- req:
	case <-quit:
		return None[Output]()
	}

	select {
	case resp := <-req.resp:
		return Some(resp)
	case <-quit:
		return None[Output]()
	}
}

// Resolve sends a response back to the requester
func (r Req[Input, Output]) Resolve(output Output) {
	r.resp <- output
}
