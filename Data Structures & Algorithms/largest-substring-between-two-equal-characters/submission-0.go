func maxLengthBetweenEqualCharacters(s string) int {
	charMap := make(map[byte]int)
	res := -1
	for i := range s {
		if j, ok := charMap[s[i]]; ok {
			c := i - j - 1
			if c > res {
				res = c
			}
			continue
		}
		charMap[s[i]] = i
	}
	return res
}