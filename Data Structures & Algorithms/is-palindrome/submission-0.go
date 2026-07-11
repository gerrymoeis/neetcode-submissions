func isAlphanum(c byte) bool {
	return (c >= 65 && c <= 90) || (c >= 97 && c <= 122) || (c >= 48 && c <= 57)
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if !isAlphanum(s[left]) {
			left++
			continue
		}
		if !isAlphanum(s[right]) {
			right--
			continue
		}
		if s[left]|' ' != s[right]|' ' {
			return false
		}
		left++
		right--
	}
	return true
}