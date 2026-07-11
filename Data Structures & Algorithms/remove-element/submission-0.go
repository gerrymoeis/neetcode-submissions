func removeElement(nums []int, val int) int {
	var k, l, r int
	r = len(nums) - 1
	for l <= r {
		if nums[l] != val {
			k++
			l++
		} else {
			nums[l], nums[r] = nums[r], nums[l]
			r--
		}
	}
	return k
}