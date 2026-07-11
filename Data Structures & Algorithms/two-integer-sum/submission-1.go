func twoSum(nums []int, target int) [2]int {
	h := make(map[int]int)
	for i, num := range nums {
		if v, ok := h[num]; ok {
			return [2]int{v, i}
		}
		h[target-num] = i
	}
	return [2]int{0, 0}
}
