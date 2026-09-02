func makeEqual(words []string) bool {
	cMap := make(map[byte]int)
	min := math.MaxInt
	for _, w := range words {
		if len(w) < min {
			min = len(w)
		}
		for i := range w {
			cMap[w[i]]++
		}
	}
	fmt.Println(cMap, min, 1 % 2)
	for _, n := range cMap {
		if n % len(words) != 0 {
			return false
		}
	}
	return true
}