package linkedlist

import (
	"errors"
)

type Node struct {
	next  *Node
	value int
}

type List struct {
	node *Node
	size int
}

func New(elements []int) *List {
	list := &List{}
	for _, v := range elements {
		list.Push(v)
	}

	return list
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	head := l.node
	if head == nil {
		l.node = &Node{value: element}
		l.size++
		return
	}
	for head.next != nil {
		head = head.next
	}

	head.next = &Node{value: element}
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.size == 0 {
		return 0, errors.New("empty list")
	}
	head := l.node
	prev := head
	for head.next != nil {
		prev = head
		head = head.next
	}

	value := head.value
	prev.next = nil
	l.size--
	return value, nil
}

func (l *List) Array() []int {
	arrRes := []int{}
	head := l.node
	for head != nil {
		arrRes = append(arrRes, head.value)
		head = head.next
	}
	return arrRes
}

func (l *List) Reverse() *List {
	curr := l.node
	var prev *Node
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.node = prev
	return l
}
