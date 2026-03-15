package fn

import (
	"testing"
	"time"
)

func TestEventDistributorBasic(t *testing.T) {
	dist := NewEventDistributor[string]()
	r1 := NewEventReceiver[string](10)
	r2 := NewEventReceiver[string](10)

	dist.Subscribe(r1)
	dist.Subscribe(r2)
	assertEqual(t, dist.NumSubscribers(), 2)

	dist.NotifySubscribers("hello")

	select {
	case v := <-r1.Events():
		assertEqual(t, v, "hello")
	case <-time.After(time.Second):
		t.Fatal("timeout r1")
	}

	select {
	case v := <-r2.Events():
		assertEqual(t, v, "hello")
	case <-time.After(time.Second):
		t.Fatal("timeout r2")
	}
}

func TestEventDistributorUnsubscribe(t *testing.T) {
	dist := NewEventDistributor[int]()
	r := NewEventReceiver[int](10)

	dist.Subscribe(r)
	dist.Unsubscribe(r)
	assertEqual(t, dist.NumSubscribers(), 0)

	dist.NotifySubscribers(42)

	select {
	case <-r.Events():
		t.Fatal("should not receive after unsubscribe")
	default:
	}
}

func TestEventReceiverStop(t *testing.T) {
	dist := NewEventDistributor[int]()
	r := NewEventReceiver[int](1)
	dist.Subscribe(r)

	r.Stop()

	// Notify should skip stopped receiver
	dist.NotifySubscribers(1)

	select {
	case <-r.Done():
	default:
		t.Fatal("Done should be closed after Stop")
	}

	// Double stop is safe
	r.Stop()
}

func TestEventDistributorFullBuffer(t *testing.T) {
	dist := NewEventDistributor[int]()
	r := NewEventReceiver[int](1)
	dist.Subscribe(r)

	// Fill buffer
	dist.NotifySubscribers(1)
	// This should be dropped (buffer full)
	dist.NotifySubscribers(2)

	val := <-r.Events()
	assertEqual(t, val, 1)

	// Channel should be empty now
	select {
	case <-r.Events():
		t.Fatal("should be empty")
	default:
	}
}
