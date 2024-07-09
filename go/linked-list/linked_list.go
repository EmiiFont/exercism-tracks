package linkedlist

import "errors"

// Define List and Node types here.
type List struct {
	head *Node
	tail *Node
}

type Node struct {
	Value interface{}
	next  *Node
	prev  *Node
}

func NewList(elements ...interface{}) *List {
	list := &List{}
	for _, val := range elements {
		list.Push(val)
	}
	return list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	node := &Node{Value: v}
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	node.next = l.head
	l.head.prev = node
	l.head = node
}

func (l *List) Push(v interface{}) {
	node := &Node{Value: v}

	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}
	node.prev = l.tail
	l.tail.next = node
	l.tail = node
}

func (l *List) Shift() (interface{}, error) {
	if l.head == nil {
		return 0, errors.New("list empty")
	}

	value := l.head.Value
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}
	return value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.tail == nil {
		return 0, errors.New("list empty")
	}

	prev := l.tail.prev
	value := l.tail.Value
	l.tail = prev
	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}
	return value, nil
}

func (l *List) Reverse() {
	if l.tail == nil {
		return
	}

	l.head, l.tail = l.tail, l.head

	node := l.head

	for node != nil {
		node.prev, node.next = node.next, node.prev
		node = node.next
	}
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
