func isMonotonic(nums []int) bool {
	var first, increasing bool
	prev := nums[0]
	for _, num := range nums[1:] {
		if num > prev {
			if !first {
				first = true
				increasing = true
			} else if !increasing {
				return false
			}
		} else if num < prev {
			if !first {
				first = true
				increasing = false
			} else if increasing {
				return false
			}
		}
		prev = num
	}
	return true
}