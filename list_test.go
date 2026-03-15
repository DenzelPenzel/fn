package fn

import "testing"

func TestListBasic(t *testing.T) {
	l := NewList[int]()
	assertEqual(t, l.Len(), 0)

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	assertEqual(t, l.Len(), 3)
	assertEqual(t, l.Front().Value, 1)
	assertEqual(t, l.Back().Value, 3)
}

func TestListPushFront(t *testing.T) {
	l := NewList[int]()
	l.PushFront(3)
	l.PushFront(2)
	l.PushFront(1)
	assertEqual(t, l.Front().Value, 1)
	assertEqual(t, l.Back().Value, 3)
}

func TestListRemove(t *testing.T) {
	l := NewList[int]()
	n := l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	val := l.Remove(n)
	assertEqual(t, val, 1)
	assertEqual(t, l.Len(), 2)
	assertEqual(t, l.Front().Value, 2)
}

func TestListInsertBeforeAfter(t *testing.T) {
	l := NewList[int]()
	n2 := l.PushBack(2)
	l.InsertBefore(1, n2)
	l.InsertAfter(3, n2)

	assertEqual(t, l.Front().Value, 1)
	assertEqual(t, l.Front().Next().Value, 2)
	assertEqual(t, l.Back().Value, 3)
}

func TestListMoveToFrontBack(t *testing.T) {
	l := NewList[int]()
	l.PushBack(1)
	l.PushBack(2)
	n3 := l.PushBack(3)

	l.MoveToFront(n3)
	assertEqual(t, l.Front().Value, 3)

	l.MoveToBack(n3)
	assertEqual(t, l.Back().Value, 3)
}

func TestListMoveBefore(t *testing.T) {
	l := NewList[int]()
	n1 := l.PushBack(1)
	l.PushBack(2)
	n3 := l.PushBack(3)

	l.MoveBefore(n3, n1)
	assertEqual(t, l.Front().Value, 3)
	assertEqual(t, l.Front().Next().Value, 1)
}

func TestListMoveAfter(t *testing.T) {
	l := NewList[int]()
	n1 := l.PushBack(1)
	l.PushBack(2)
	n3 := l.PushBack(3)

	l.MoveAfter(n1, n3)
	assertEqual(t, l.Back().Value, 1)
}

func TestListPushBackList(t *testing.T) {
	l1 := NewList[int]()
	l1.PushBack(1)
	l1.PushBack(2)

	l2 := NewList[int]()
	l2.PushBack(3)
	l2.PushBack(4)

	l1.PushBackList(l2)
	assertEqual(t, l1.Len(), 4)
	assertEqual(t, l1.Back().Value, 4)
}

func TestListPushFrontList(t *testing.T) {
	l1 := NewList[int]()
	l1.PushBack(3)
	l1.PushBack(4)

	l2 := NewList[int]()
	l2.PushBack(1)
	l2.PushBack(2)

	l1.PushFrontList(l2)
	assertEqual(t, l1.Len(), 4)
	assertEqual(t, l1.Front().Value, 1)
}

func TestListFilter(t *testing.T) {
	l := NewList[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)

	even := l.Filter(func(i int) bool { return i%2 == 0 })
	assertEqual(t, even.Len(), 2)
	assertEqual(t, even.Front().Value, 2)
	assertEqual(t, even.Back().Value, 4)
}

func TestListNodeNavigation(t *testing.T) {
	l := NewList[int]()
	n1 := l.PushBack(1)
	n2 := l.PushBack(2)
	n3 := l.PushBack(3)

	// Next/Prev
	assertEqual(t, n1.Next(), n2)
	assertEqual(t, n2.Next(), n3)
	assertTrue(t, n3.Next() == nil)

	assertTrue(t, n1.Prev() == nil)
	assertEqual(t, n2.Prev(), n1)
	assertEqual(t, n3.Prev(), n2)
}

func TestListEmpty(t *testing.T) {
	l := NewList[int]()
	assertTrue(t, l.Front() == nil)
	assertTrue(t, l.Back() == nil)
}

func TestListInit(t *testing.T) {
	l := NewList[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.Init()
	assertEqual(t, l.Len(), 0)
	assertTrue(t, l.Front() == nil)
}
