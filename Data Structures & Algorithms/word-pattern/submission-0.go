func wordPattern(pattern string, s string) bool {
    p_to_s := make(map[byte]string)
	s_to_p := make(map[string]byte)
	var i, k, offset int
	for j, c := range s {
		if c == ' ' || j == len(s)-1 {
			if j == len(s)-1 {
				offset = 1
			} else {
				offset = 0
			}
			if k >= len(pattern) {
				return false
			}
			word := s[i:j+offset]
			p := pattern[k]
			if curr, ok := p_to_s[p]; ok && curr != word {
				return false
			} else if curr, ok := s_to_p[word]; ok && curr != p {
				return false
			}
			p_to_s[p] = word
			s_to_p[word] = p
			i = j+offset+1
			k++
		}
	}
	if k < len(pattern) {
		return false
	}
	return true
}