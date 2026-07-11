func removeDuplicates(nums []int) int {
	l := 0
	for _, num := range nums[1:] {
		if num != nums[l] {
			l++
			nums[l] = num
		}
	}
	return l + 1
}