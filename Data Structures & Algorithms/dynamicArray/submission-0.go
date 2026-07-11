type DynamicArray struct {
	data     []int
	size     int
	capacity int
}

func NewDynamicArray(capacity int) *DynamicArray {
	return &DynamicArray{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (da *DynamicArray) Get(i int) int {
	if i < 0 || i >= da.size {
		return 0
	}
	return da.data[i]
}

func (da *DynamicArray) Set(i int, n int) {
	if i < 0 || i >= da.size {
		return
	}
	da.data[i] = n
}

func (da *DynamicArray) Pushback(n int) {
	if da.size == da.capacity {
		da.resize()
	}
	da.data[da.size] = n
	da.size++
}

func (da *DynamicArray) Popback() int {
	if da.size == 0 {
		return 0
	}
	last := da.data[da.size-1]
	da.data[da.size-1] = 0
	da.size--
	return last
}

func (da *DynamicArray) resize() {
	da.capacity *= 2
	newData := make([]int, da.capacity)
	copy(newData[:da.size], da.data[:da.size])
	da.data = newData
}

func (da *DynamicArray) GetSize() int {
	return da.size
}

func (da *DynamicArray) GetCapacity() int {
	return da.capacity
}
