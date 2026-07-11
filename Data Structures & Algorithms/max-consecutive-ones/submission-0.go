func findMaxConsecutiveOnes(nums []int) int {
	c := 0
	r := 0
	for _, num := range nums {
		if num == 1 {
			c++
		} else {
			c = 0
		}
		if c > r {
			r = c
		}
	}
	return r
}
