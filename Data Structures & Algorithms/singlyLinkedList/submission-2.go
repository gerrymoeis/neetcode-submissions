type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) GetNode(index int) *Node {
	node := ll.Head
	if index == -1 {
		for node.Next != nil {
			node = node.Next
		}
	} else {
		for i := 0; i < index; i++ {
			if node == nil {
				break
			}
			node = node.Next
		}
	}
	return node
}

func (ll *LinkedList) Get(index int) int {
	if index < 0 || ll.Head == nil {
		return -1
	}
	node := ll.GetNode(index)
	if node == nil {
		return -1
	}
	return node.Val
}

func (ll *LinkedList) InsertHead(val int) {
	if ll.Head == nil {
		ll.Head = &Node{Val: val}
		return
	}
	newNode := &Node{Val: val, Next: ll.Head}
	ll.Head = newNode
}

func (ll *LinkedList) InsertTail(val int) {
	if ll.Head == nil {
		ll.Head = &Node{Val: val}
		return
	}
	node := ll.Head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = &Node{Val: val}
}

func (ll *LinkedList) Remove(index int) bool {
	if index < 0 || ll.Head == nil {
		return false
	}
	if index == 0 {
		if ll.Head.Next == nil {
			ll.Head = nil
		} else {
			ll.Head = ll.Head.Next
		}
	} else {
		prev := ll.GetNode(index - 1)
		if prev == nil || prev.Next == nil {
			return false
		}
		prev.Next = prev.Next.Next
	}
	return true
}

func (ll *LinkedList) GetValues() []int {
	var values []int
	for node := ll.Head; node != nil; node = node.Next {
		values = append(values, node.Val)
	}
	return values
}