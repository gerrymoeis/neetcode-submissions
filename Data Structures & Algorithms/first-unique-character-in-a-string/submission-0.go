func firstUniqChar(s string) int {
	charMap := make(map[rune]*[2]int)
	for i, c := range s {
		if _, ok := charMap[c]; ok {
			charMap[c][1] = 1
		} else {
			charMap[c] = &[2]int{i, 0}
		}
	}
	min := math.MaxInt
	for _, val := range charMap {
		if val[1] == 0 {
			if val[0] < min {
				min = val[0]
			}
		}
	}
	if min == math.MaxInt {
		min = -1
	}
	return min
}
