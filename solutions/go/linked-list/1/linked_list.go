package linkedlist

// Define List and Node types here.
type List struct {
	Head *Node
}
type Node struct {
	Value interface{}
	next  *Node
	prev  *Node
}

// Note: The tests expect Node type to include an exported field with name Value to pass.

func NewList(elements ...interface{}) *List {
	return &List{
		Head: &Node{
			Value: elements,
			next:  nil,
			prev:  nil,
		},
	}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	panic("Please implement the Unshift function")
}

func (l *List) Push(v interface{}) {
	panic("Please implement the Push function")
}

func (l *List) Shift() (interface{}, error) {
	panic("Please implement the Shift function")
}

func (l *List) Pop() (interface{}, error) {
	panic("Please implement the Pop function")
}

func (l *List) Reverse() {
	panic("Please implement the Reverse function")
}

func (l *List) First() *Node {
	return l.Head
}

func (l *List) Last() *Node {
	panic("Please implement the Last function")
}
