type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type MyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) GetNode(index int) *Node {
	if this.Head == nil && this.Tail == nil {
		return nil
	}
	var node *Node
	if index <= this.Length/2 {
		node = this.Head
		for i := 0; i < index; i++ {
			node = node.Next
		}
	} else {
		node = this.Tail
		for i := this.Length - 1; i > index; i-- {
			node = node.Prev
		}
	}
	return node
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.Length-1 || (this.Head == nil && this.Tail == nil) {
		return -1
	}
	return this.GetNode(index).Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	if this.Head == nil && this.Tail == nil {
		this.Head = &Node{Val: val}
		this.Tail = this.Head
	} else {
		oldHead := this.Head
		this.Head = &Node{Val: val, Next: oldHead}
		oldHead.Prev = this.Head
	}
	this.Length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.Tail == nil && this.Head == nil {
		this.Tail = &Node{Val: val}
		this.Head = this.Tail
	} else {
		oldTail := this.Tail
		this.Tail = &Node{Val: val, Prev: oldTail}
		oldTail.Next = this.Tail
	}
	this.Length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Length {
		return
	}
	switch index {
	case 0:
		this.AddAtHead(val)
	case this.Length:
		this.AddAtTail(val)
	default:
		node := this.GetNode(index - 1)
		if node == nil {
			return
		}
		oldNext := node.Next
		node.Next = &Node{Val: val, Next: oldNext, Prev: node}
		if oldNext != nil {
			oldNext.Prev = node.Next
		}
		this.Length++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.Length-1 || (this.Head == nil && this.Tail == nil) {
		return
	}
	switch index {
	case 0:
		this.Head = this.Head.Next
		if this.Head != nil {
			this.Head.Prev = nil
		}
		if this.Length == 1 {
			this.Tail = nil
		}
	case this.Length - 1:
		this.Tail = this.Tail.Prev
		if this.Tail != nil {
			this.Tail.Next = nil
		}
		if this.Length == 1 {
			this.Head = nil
		}
	default:
		node := this.GetNode(index)
		if node.Prev != nil {
			node.Prev.Next = node.Next
		}
		if node.Next != nil {
			node.Next.Prev = node.Prev
		}
	}
	this.Length--
}