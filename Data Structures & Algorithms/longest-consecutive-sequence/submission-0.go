func longestConsecutive(nums []int) int {
	maxLength := 0
	diffMaps := make(map[int]int)
	starters := make(map[int]int)
	for _, num := range nums {
		diffMaps[num+1] = num
	}
	for _, val := range diffMaps {
		if _, ok := diffMaps[val]; !ok {
			starters[val] = val+1
		}
	}
	for _, val := range starters {
		count := 0
		for {
			_, ok := diffMaps[val]
			if !ok {
				break
			}
			val++
			count++
		}
		if count > maxLength {
			maxLength = count
		}
	}
	return maxLength
}