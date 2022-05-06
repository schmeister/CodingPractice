package linkedlist

import "errors"


var ErrEmptyList = errors.New("empty list")

// Define List and Node types here.
// Note: The tests expect Node type to include an
// exported field with name Value to pass.
type Node struct {
	Value interface{}
	prev  *Node
	next  *Node
}
type List struct {
	head *Node
	tail *Node
}

func NewList(args ...interface{}) *List {
	list := &List{}
	for _, a := range args {
		list.PushBack(a)
	}
	return list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) PushFront(v interface{}) {
	new := &Node{Value: v}
	// New node becomes head
	new.next = l.head

	// if list is not empty, pointback to new
	if l.head != nil {
		l.head.prev = new
	}

	// Set head to new
	l.head = new

	// If list was empty, also set tail to new
	if l.tail == nil {
		l.tail = new
	}
}

func (l *List) PushBack(v interface{}) {
	new := &Node{Value: v}
	// New node becomes tail
	new.prev = l.tail

	// if list is not empty, pointforward to new
	if l.tail != nil {
		l.tail.next = new
	}

	// Set tail to new
	l.tail = new

	// If list was empty, also set head to new
	if l.head == nil {
		l.head = new
	}
}

func (l *List) PopFront() (interface{}, error) {
	if l.head == nil {
        return nil, ErrEmptyList
    }
	n := l.head
    l.head = n.next
    if l.head == nil {
        l.tail = nil
    } else {
    	l.head.prev = nil
    }
	n.next = nil
    return n.Value, nil
}

func (l *List) PopBack() (interface{}, error) {
	if l.tail == nil {
        return nil, ErrEmptyList
    }
	n := l.tail
    l.tail = n.prev
    if l.tail == nil {
        l.head = nil
    } else {
    	l.tail.next = nil
    }
	n.prev = nil
    return n.Value, nil
}

func (l *List) Reverse() {
	n := l.head
	for n != nil {
		n.next, n.prev = n.prev, n.next
		n = n.prev
	}
	l.head, l.tail = l.tail, l.head
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
