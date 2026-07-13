func longestConsecutive(nums []int) int {
	maxLength := 0
	starters := make(map[int]bool)
	for _, num := range nums {
		starters[num] = true
	}
	for num := range starters {
		if _, ok := starters[num-1]; !ok {
			count := 0
			for {
				_, ok := starters[num]
				if !ok {
					break
				}
				num++
				count++
			}
			if count > maxLength {
				maxLength = count
			}
		}
	}
	return maxLength
}