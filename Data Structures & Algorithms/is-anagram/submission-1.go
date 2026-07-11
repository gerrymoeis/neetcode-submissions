func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	l := [26]int{}
	for i:=0; i<len(s); i++ {
		l[s[i]-'a']++
		l[t[i]-'a']--
	}
	for _, v := range l {
		if (v != 0) {
			return false
		}
	}
	return true
}
