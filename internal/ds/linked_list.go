package ds

type Node struct {
	Val      int
	Next     *Node
	Previous *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Insert sets the desired insertion as the head of the linked list
func (l *LinkedList) Insert(num int) {
	newNode := &Node{
		Val:      num,
		Next:     l.head,
		Previous: nil,
	}

	if l.head != nil {
		l.head.Previous = newNode
	}

	l.head = newNode
}

func (l *LinkedList) Delete(needle int) {
	x := l.Search(needle)

	if x.Previous != nil {
		x.Previous.Next = x.Next
	} else {
		l.head = x.Next
	}
	if x.Next != nil {
		x.Next.Previous = x.Previous
	}
}

// Search iterates through the list and returns the corresponding node, if not found, nil is returned
func (l *LinkedList) Search(needle int) *Node {
	curr := l.head

	for curr != nil && curr.Val != needle {
		curr = curr.Next
	}
	return curr
}

func (l *LinkedList) FromArray(arr []int) {
	for _, item := range arr {
		l.Insert(item)
	}
}

// ToArray transforms linked list to array keeping the order
func (l *LinkedList) ToArray() []int {
	result := make([]int, l.len)

	ptr := l.head
	for ptr != nil {
		result = append(result, ptr.Val)
		ptr = ptr.Next
	}

	return result
}
