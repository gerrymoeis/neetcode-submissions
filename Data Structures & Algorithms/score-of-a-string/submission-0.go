func scoreOfString(s string) int {
	score := 0
	for i := 1; i < len(s); i++ {
		adj := int(s[i]) - int(s[i-1])
		if adj < 0 {
			adj = -adj
		}
		score += adj
	}
	return score
}
