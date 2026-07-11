func findMaxConsecutiveOnes(nums []int) int {
	var c, r int
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
