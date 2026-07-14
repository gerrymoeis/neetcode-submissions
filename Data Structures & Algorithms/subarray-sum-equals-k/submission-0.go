func subarraySum(nums []int, k int) int {
	count := 0
	sums := make(map[int]int)
	sums[0] = 1
	total := 0
	for _, num := range nums {
		total += num
		count += sums[total-k]
		sums[total]++
	}
	return count
}