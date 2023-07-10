package list

// IList is an interface for generic linked list
type IList[T any] interface {
	Len() int
	Append(T)
	First() *T
	Next() *T
	Remove() *T
}

// Node is a node of generic linked list of type T
type Node[T any] struct {
	next  *Node[T]
	value T
}

// List is generic linked list of type T
type List[T any] struct {
	head    Node[T]
	current *Node[T]
}

// verify interface
var _ IList[int] = new(List[int])

// New creates a new linked list
func New[T any]() *List[T] {
	l := &List[T]{
		head: Node[T]{},
	}
	return l
}

// Len returns the number of elements in the linked list l.
func (l *List[T]) Len() int {
	len := 0
	for ll := l.head.next; ll != nil; ll = ll.next {
		len++
	}
	return len
}

// Append adds an element to the end of the linked list
func (l *List[T]) Append(el T) {
	if l.head.next == nil {
		l.head.next = &Node[T]{value: el}
		l.current = l.head.next
		return
	}
	for ll := l.head.next; ll != nil; ll = ll.next {
		if ll.next == nil {
			ll.next = &Node[T]{value: el}
			l.current = l.head.next
			return
		}
	}
}

func (l *List[T]) First() *T {
	if l.head.next == nil {
		return nil
	}
	l.current = l.head.next
	return &l.current.value
}

func (l *List[T]) Next() *T {
	if l.current == nil {
		return l.First()
	}
	l.current = l.current.next
	if l.current == nil {
		return nil
	}
	return &l.current.value
}

// Remove removes current element from the linked list
func (l *List[T]) Remove() *T {
	if l.current == nil {
		return nil
	}
	res := l.current.value
	//      V
	// C -> D -> E ...
	// C -> E  ...
	// replace current node with next node
	// delete next node
	if l.current.next != nil {
		l.current.value = l.current.next.value
		l.current.next = l.current.next.next
	} else {
		l.head.next = nil
	}
	return &res
}
