func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	l := make(map[byte]int)
	for i:=0; i<len(s); i++ {
		l[s[i]]++
		l[t[i]]--
	}
	for _, v := range l {
		if (v < 0) {
			return false
		}
	}
	return true
}
