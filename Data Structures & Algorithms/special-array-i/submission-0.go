func isArraySpecial(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return true
	}
	prev := nums[0]
	for _, num := range nums[1:] {
		if (num+prev)%2 == 0 {
			return false
		}
		prev = num
	}
	return true
}