type Node struct {
	Val  int
	Next *Node
}

type MyStack struct {
	top *Node
}

func Constructor() *MyStack {
	return &MyStack{}
}

func (s *MyStack) Push(x int) {
	s.top = &Node{Val: x, Next: s.top}
}

func (s *MyStack) Pop() int {
	if s.top == nil {
		return -1
	}
	val := s.top.Val
	s.top = s.top.Next
	return val
}

func (s *MyStack) Top() int {
	if s.top == nil {
		return -1
	}
	return s.top.Val
}

func (s *MyStack) Empty() bool {
	return s.top == nil
}