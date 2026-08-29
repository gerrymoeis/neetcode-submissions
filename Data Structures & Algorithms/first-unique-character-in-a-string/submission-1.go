func firstUniqChar(s string) int {
	n := len(s)
	charMap := make(map[rune]int)
	for i, c := range s {
		if _, ok := charMap[c]; ok {
			charMap[c] = n
		} else {
			charMap[c] = i
		}
	}
	min := n
	for _, i := range charMap {
		if i < min {
			min = i
		}
	}
	if min == n {
		min = -1
	}
	return min
}
