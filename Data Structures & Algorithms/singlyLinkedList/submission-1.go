type Node struct {
	Val  int
	Next *Node
	Prev *Node
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
	var node *Node
	if index < ll.Len/2 {
		node = ll.Head
		for i := 0; i < index; i++ {
			node = node.Next
		}
	} else {
		node = ll.Tail
		for i := ll.Len - 1; i > index; i-- {
			node = node.Prev
		}
	}
	return node.Val
}

func (ll *LinkedList) InsertHead(val int) {
	node := &Node{Val: val}
	node.Next = ll.Head
	if ll.Head != nil {
		ll.Head.Prev = node
	}
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
		node.Prev = ll.Tail
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
		if ll.Head != nil {
			ll.Head.Prev = nil
		}
	case ll.Len - 1:
		node = ll.Tail
		ll.Tail = node.Next
	default:
		for i := 0; i < index-1; i++ {
			node = node.Next
		}
		node.Next = node.Next.Next
		if node.Next != nil {
			node.Next.Prev = node
		}
	}
	ll.Len--
	if ll.Len == 0 {
		ll.Tail = nil
	}
	return true
}

func (ll *LinkedList) GetValues() []int {
	values := make([]int, ll.Len)
	node_1 := ll.Head
	node_2 := ll.Tail
	for i := 0; i < ll.Len/2; i++ {
		values[i] = node_1.Val
		node_1 = node_1.Next
		values[ll.Len-1-i] = node_2.Val
		node_2 = node_2.Prev
	}
	if ll.Len%2 == 1 {
		values[ll.Len/2] = node_1.Val
	}
	return values
}