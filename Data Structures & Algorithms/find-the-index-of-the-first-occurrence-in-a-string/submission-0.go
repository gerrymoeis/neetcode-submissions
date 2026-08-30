func strStr(haystack string, needle string) int {
	n := len(needle)
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] && i+n <= len(haystack) && haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}
