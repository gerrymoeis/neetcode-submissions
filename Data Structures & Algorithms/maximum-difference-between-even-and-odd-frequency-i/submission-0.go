func maxDifference(s string) int {
	maxOdd, minEven := 0, math.MaxInt
	cFreq := make(map[byte]int)
	for i := range s {
		cFreq[s[i]]++
	}
	for _, freq := range cFreq {
		if freq % 2 != 0 {
			if freq > maxOdd {
				maxOdd = freq
			}
		} else {
			if freq < minEven {
				minEven = freq
			}
		}
	}
	return maxOdd - minEven
}
