func longestCommonPrefix(strs []string) string {
	i := 0
	for {
		if i >= len(strs[0]) {
			return strs[0][:i]
		}
		first := strs[0][i]
		for _, s := range strs {
			if i >= len(s) {
				return strs[0][:i]
			}
			if first != s[i] {
				return strs[0][:i]
			}
		}
		i++
	}
	return ""
}
