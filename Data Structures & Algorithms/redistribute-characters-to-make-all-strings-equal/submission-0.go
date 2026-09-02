func makeEqual(words []string) bool {
	cMap := make(map[byte]int)
	for _, w := range words {
		for i := range w {
			cMap[w[i]]++
		}
	}
	for _, n := range cMap {
		if n%len(words) != 0 {
			return false
		}
	}
	return true
}