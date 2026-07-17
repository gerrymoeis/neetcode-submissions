type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Len  int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Get(index int) int {
	if index < 0 || index >= ll.Len {
		return -1
	}
	node := ll.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node.Val
}

func (ll *LinkedList) InsertHead(val int) {
	node := &Node{Val: val}
	node.Next = ll.Head
	ll.Head = node
	ll.Len++
	if ll.Tail == nil {
		ll.Tail = node
	}
}

func (ll *LinkedList) InsertTail(val int) {
	node := &Node{Val: val}
	if ll.Tail != nil {
		ll.Tail.Next = node
	}
	ll.Tail = node
	ll.Len++
	if ll.Head == nil {
		ll.Head = node
	}
}

func (ll *LinkedList) Remove(index int) bool {
	if index < 0 || index >= ll.Len {
		return false
	}
	node := ll.Head
	switch index {
	case 0:
		ll.Head = node.Next
	case ll.Len - 1:
		node = ll.Tail
		ll.Tail = node.Next
	default:
		for i := 0; i < index-1; i++ {
			node = node.Next
		}
		node.Next = node.Next.Next
	}
	ll.Len--
	if ll.Len == 0 {
		ll.Tail = nil
	}
	return true
}

func (ll *LinkedList) GetValues() []int {
	values := make([]int, ll.Len)
	node := ll.Head
	for i := 0; i < ll.Len; i++ {
		values[i] = node.Val
		node = node.Next
	}
	return values
}
