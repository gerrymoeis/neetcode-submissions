type MyStack struct {
	items []int
}

func Constructor() *MyStack {
	return &MyStack{}
}

func (s *MyStack) Push(x int) {
	s.items = append(s.items, x)
}

func (s *MyStack) Pop() int {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *MyStack) Top() int {
	return s.items[len(s.items)-1]
}
func (s *MyStack) Empty() bool {
	return len(s.items) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
