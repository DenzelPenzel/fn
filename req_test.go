package fn

import "testing"

func TestReqDispatchResolve(t *testing.T) {
	ch := make(chan Req[string, int])
	quit := make(chan struct{})

	// Server goroutine
	go func() {
		req := <-ch
		assertEqual(t, req.Input, "hello")
		req.Resolve(42)
	}()

	req := NewReq[string, int]("hello")
	resp := Dispatch(req, ch, quit)
	assertEqual(t, resp.UnsafeFromSome(), 42)
}

func TestReqDispatchQuit(t *testing.T) {
	ch := make(chan Req[string, int])
	quit := make(chan struct{})
	close(quit)

	req := NewReq[string, int]("hello")
	resp := Dispatch(req, ch, quit)
	assertTrue(t, resp.IsNone())
}
