func maxAscendingSum(nums []int) int {
    var prev, sum, max int
	for _, num := range nums {
		if num <= prev {
			sum = 0
		}
		sum += num
		prev = num
		if sum > max {
			max = sum
		}
	}
	return max
}