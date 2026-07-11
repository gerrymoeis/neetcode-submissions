func findMaxConsecutiveOnes(nums []int) int {
	c := 0
	r := 0
	for _, num := range nums {
		if num == 1 {
			c++
			if c > r {
				r = c
			}
		} else {
			c = 0
		}
	}
	return r
}
