func isIsomorphic(s string, t string) bool {
	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if (sMap[s[i]] != 0 && sMap[s[i]] != t[i]) ||
			(tMap[t[i]] != 0 && tMap[t[i]] != s[i]) {
			return false
		}
		sMap[s[i]] = t[i]
		tMap[t[i]] = s[i]
	}
	return true
}
