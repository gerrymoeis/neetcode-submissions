func wordPattern(pattern string, s string) bool {
    p_to_s := make(map[byte]string)
	s_to_p := make(map[string]byte)
	var i, k, offset int
	for j, c := range s {
		if c == ' ' || j == len(s)-1 {
			if j == len(s)-1 {
				fmt.Println("ancok")
				offset = 1
			} else {
				fmt.Println("kontol")
				offset = 0
			}
			if k >= len(pattern) {
				return false
			}
			word := s[i:j+offset]
			p := pattern[k]
			if curr, ok := p_to_s[p]; ok && curr != word {
				fmt.Println(curr, word, len(curr), len(word))
				return false
			} else if curr, ok := s_to_p[word]; ok && curr != p {
				fmt.Println(curr, p, len(word))
				return false
			}
			p_to_s[p] = word
			s_to_p[word] = p
			i = j+offset+1
			k++
		}
	}
	fmt.Println(p_to_s)
	fmt.Println(s_to_p)
	fmt.Println(k, len(pattern))
	if k < len(pattern) {
		return false
	}
	// for k, v := range p_to_s {
	// 	fmt.Print(k)
	// 	fmt.Print(v)
	// }
	// fmt.Println()
	// for k, v := range s_to_p {
	// 	fmt.Print(k)
	// 	fmt.Print(v)
	// }
	// fmt.Println()
	// fmt.Println(string(p_to_s[97][2]), string(s_to_p["dog"]))
	return true
}