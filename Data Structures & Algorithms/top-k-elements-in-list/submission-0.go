func topKFrequent(nums []int, k int) []int {
	result := make([]int, k)
	numHash := make(map[int]int)
	freqs := make([][]int, len(nums)+1)
	for _, num := range nums {
		numHash[num]++
	}
	for num, freq := range numHash {
		freqs[freq] = append(freqs[freq], num)
	}
	for i := len(freqs) - 1; i >= 0; i-- {
		if len(freqs[i]) != 0 {
			for _, num := range freqs[i] {
				result[k-1] = num
				k--
				if k <= 0 {
					return result
				}
			}
		}
	}
	return result
}