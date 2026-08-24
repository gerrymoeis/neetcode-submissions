func numIdenticalPairs(nums []int) int {
    count := 0
	numsMap := make(map[int]int)
	for _, num := range nums {
		if _, ok := numsMap[num]; ok {
			count += numsMap[num]
		}
		numsMap[num]++
	}
	return count
}