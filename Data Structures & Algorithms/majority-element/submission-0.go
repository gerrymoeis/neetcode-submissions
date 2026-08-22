func majorityElement(nums []int) int {
    var majority, large int
	countNums := make(map[int]int)
	for _, num := range nums {
		countNums[num]++
		if countNums[num] > large {
			large = countNums[num]
			majority = num
		}
	}
	return majority
}
