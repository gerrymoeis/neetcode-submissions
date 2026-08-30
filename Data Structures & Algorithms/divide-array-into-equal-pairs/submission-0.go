func divideArray(nums []int) bool {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	for _, count := range numMap {
		if count % 2 != 0 {
			return false
		}
	}
	return true
}