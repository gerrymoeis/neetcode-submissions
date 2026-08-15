func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	i := 0
	for _, c := range t {
		if i+1 == len(s) {
			return true
		}
		if s[i] == byte(c) {
			i++
		}
	}
	return false
}
