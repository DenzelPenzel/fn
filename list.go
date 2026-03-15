package fn

// Node is an element of a doubly-linked list
type Node[A any] struct {
	Value      A
	next, prev *Node[A]
	list       *List[A]
}

// Next returns the next node or nil
func (n *Node[A]) Next() *Node[A] {
	if n.list != nil && n.next != &n.list.root {
		return n.next
	}
	return nil
}

// Prev returns the previous node or nil
func (n *Node[A]) Prev() *Node[A] {
	if n.list != nil && n.prev != &n.list.root {
		return n.prev
	}
	return nil
}

// List is a generic doubly-linked list
type List[A any] struct {
	root Node[A]
	len  int
}

// Init initializes or clears the list
func (l *List[A]) Init() *List[A] {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// NewList creates a new empty list
func NewList[A any]() *List[A] {
	return new(List[A]).Init()
}

// Len returns the number of elements
func (l *List[A]) Len() int {
	return l.len
}

// Front returns the first node or nil
func (l *List[A]) Front() *Node[A] {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back returns the last node or nil
func (l *List[A]) Back() *Node[A] {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// insert inserts n after at, increments length, and returns n
func (l *List[A]) insert(n, at *Node[A]) *Node[A] {
	n.prev = at
	n.next = at.next
	n.prev.next = n
	n.next.prev = n
	n.list = l
	l.len++
	return n
}

// insertValue is a convenience wrapper for insert
func (l *List[A]) insertValue(v A, at *Node[A]) *Node[A] {
	return l.insert(&Node[A]{Value: v}, at)
}

// remove unlinks n from the list
func (l *List[A]) remove(n *Node[A]) {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil
	n.prev = nil
	n.list = nil
	l.len--
}

// move moves n to after at
func (l *List[A]) move(n, at *Node[A]) {
	if n == at {
		return
	}
	n.prev.next = n.next
	n.next.prev = n.prev
	n.prev = at
	n.next = at.next
	n.prev.next = n
	n.next.prev = n
}

// PushFront inserts a value at the front of the list
func (l *List[A]) PushFront(v A) *Node[A] {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack inserts a value at the back of the list
func (l *List[A]) PushBack(v A) *Node[A] {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// InsertBefore inserts a value before mark
func (l *List[A]) InsertBefore(v A, mark *Node[A]) *Node[A] {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}

// InsertAfter inserts a value after mark
func (l *List[A]) InsertAfter(v A, mark *Node[A]) *Node[A] {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
}

// Remove removes n from the list and returns its value
func (l *List[A]) Remove(n *Node[A]) A {
	if n.list == l {
		l.remove(n)
	}
	return n.Value
}

// MoveToFront moves n to the front of the list
func (l *List[A]) MoveToFront(n *Node[A]) {
	if n.list != l || l.root.next == n {
		return
	}
	l.move(n, &l.root)
}

// MoveToBack moves n to the back of the list
func (l *List[A]) MoveToBack(n *Node[A]) {
	if n.list != l || l.root.prev == n {
		return
	}
	l.move(n, l.root.prev)
}

// MoveBefore moves n before mark
func (l *List[A]) MoveBefore(n, mark *Node[A]) {
	if n.list != l || n == mark || mark.list != l {
		return
	}
	l.move(n, mark.prev)
}

// MoveAfter moves n after mark
func (l *List[A]) MoveAfter(n, mark *Node[A]) {
	if n.list != l || n == mark || mark.list != l {
		return
	}
	l.move(n, mark)
}

// PushBackList appends all elements of other to the back of l. The other list
// is not modified
func (l *List[A]) PushBackList(other *List[A]) {
	l.lazyInit()
	for n := other.Front(); n != nil; n = n.Next() {
		l.insertValue(n.Value, l.root.prev)
	}
}

// PushFrontList prepends all elements of other to the front of l
func (l *List[A]) PushFrontList(other *List[A]) {
	l.lazyInit()
	for i, n := other.Len(), other.Back(); i > 0; i, n = i-1, n.Prev() {
		l.insertValue(n.Value, &l.root)
	}
}

// Filter returns a new list with only the elements matching pred
func (l *List[A]) Filter(pred Pred[A]) *List[A] {
	result := NewList[A]()
	for n := l.Front(); n != nil; n = n.Next() {
		if pred(n.Value) {
			result.PushBack(n.Value)
		}
	}
	return result
}

// lazyInit initializes the list if it hasn't been yet
func (l *List[A]) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}
