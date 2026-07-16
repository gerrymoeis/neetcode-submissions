type MinStack struct {
	data []int
	min  []int
}

func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0),
		min:  make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	if len(this.min) == 0 || val <= this.GetMin() {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.data) == 0 {
		return
	}
	if this.Top() == this.GetMin() {
		this.min = this.min[:len(this.min)-1]
	}
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return 0
	}
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 0
	}
	return this.min[len(this.min)-1]
}