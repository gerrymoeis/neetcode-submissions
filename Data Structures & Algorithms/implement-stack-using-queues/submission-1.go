type Queue struct {
	Val  int
	Prev *Queue
}

type MyStack struct {
	top *Queue
	len int
}

func Constructor() *MyStack {
	return &MyStack{}
}

func (s *MyStack) Push(x int) {
	node := &Queue{Val: x}
	if s.len == 0 {
		s.top = node
	} else {
		node.Prev = s.top
		s.top = node
	}
	s.len++
}

func (s *MyStack) Pop() int {
	if s.len == 0 {
		return 0
	}
	val := s.top.Val
	s.top = s.top.Prev
	s.len--
	return val
}

func (s *MyStack) Top() int {
	if s.len == 0 {
		return 0
	}
	return s.top.Val
}

func (s *MyStack) Empty() bool {
	return s.len == 0
}