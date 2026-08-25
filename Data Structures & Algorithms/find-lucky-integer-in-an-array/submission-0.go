func findLucky(arr []int) int {
	freqs := make(map[int]int)
	for _, n := range arr {
		freqs[n]++
	}
	largest := -1
	for n, freq := range freqs {
		if n == freq && n > largest {
			largest = n
		}
	}
	return largest
}