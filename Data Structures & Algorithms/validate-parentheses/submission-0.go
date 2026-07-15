func isValid(s string) bool {
	n := len(s)
	if n == 1 {
		return false
	}
	if n == 0 {
		return true
	}
	for i := 0; i <= n/2; i++ {
		l, r := s[i], s[n-1-i]
		if l == 40 {
			l++
		} else {
			l += 2
		}
		if r == 41 {
			r--
		} else {
			r -= 2
		}
		if l == s[i+1] && r == s[n-2-i] {
			i++
		} else if l != s[n-1-i] {
			return false
		}
	}
	return true
}
