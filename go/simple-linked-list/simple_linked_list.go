package linkedlist

// Define the List and Element types here.
type Node struct {
	next  *Node
	value int
}

type List struct {
	node Node
	size int
}

func New(elements []int) *List {
	node := Node{}
	for _, val := range elements {
		node.next = &Node{value: val}
		node = *node.next

	}
	return &List{node: node, size: len(elements)}
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	l.size++
	dummy := l.node
	for dummy.value != 0 {
		dummy = *dummy.next
	}

	dummy.next = &Node{value: element}
}

func (l *List) Pop() (int, error) {
	l.size--
	panic("Please implement the Pop function")
}

func (l *List) Array() []int {
	panic("Please implement the Array function")
}

func (l *List) Reverse() *List {
	panic("Please implement the Reverse function")
}
