func productExceptSelf(nums []int) []int {
	product := 1
	isZero := false
	for _, num := range nums {
		if num == 0 && !isZero {
			isZero = true
			continue
		}
		product *= num
	}
	for i, num := range nums {
		if num != 0 && isZero {
			nums[i] = 0
		} else if num == 0 {
			nums[i] = product
		} else {
			nums[i] = product / num
		}
	}
	return nums
}
