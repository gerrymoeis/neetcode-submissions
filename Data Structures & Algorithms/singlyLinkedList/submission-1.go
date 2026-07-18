const Uninitialized = math.MinInt

type LinkedList struct {
	Val  int
	Next *LinkedList
}

func NewLinkedList() *LinkedList {
	return &LinkedList{Val: Uninitialized}
}

func (ll *LinkedList) GetNode(index int) *LinkedList {
	node := ll
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
	if index < 0 || ll.Val == Uninitialized {
		return -1
	}
	node := ll.GetNode(index)
	if node == nil {
		return -1
	}
	return node.Val
}

func (ll *LinkedList) InsertHead(val int) {
	if ll.Val == Uninitialized {
		ll.Val = val
		return
	}
	newNode := &LinkedList{Val: ll.Val, Next: ll.Next}
	ll.Val = val
	ll.Next = newNode
}

func (ll *LinkedList) InsertTail(val int) {
	if ll.Val == Uninitialized {
		ll.Val = val
		return
	}
	newNode := &LinkedList{Val: val}
	node := ll.GetNode(-1)
	node.Next = newNode
}

func (ll *LinkedList) Remove(index int) bool {
	if index < 0 || ll.Val == Uninitialized {
		return false
	}
	if index == 0 {
		if ll.Next == nil {
			ll.Val = Uninitialized
		} else {
			ll.Val = ll.Next.Val
			ll.Next = ll.Next.Next
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
	if ll.Val == Uninitialized {
		return nil
	}
	var values []int
	for node := ll; node != nil; node = node.Next {
		values = append(values, node.Val)
	}
	return values
}