func makeKey(s string) [26]byte {
	c := [26]byte{}
	for i := 0; i < len(s); i++ {
		c[s[i]-'a']++
	}
	return c
}

func groupAnagrams(strs []string) [][]string {
	r := make(map[[26]byte][]string)
	for _, str := range strs {
		k := makeKey(str)
		if _, ok := r[k]; ok {
			r[k] = append(r[k], str)
		} else {
			r[k] = []string{str}
		}
	}
	result := make([][]string, len(r))
	i := 0
	for _, v := range r {
		result[i] = v
		i++
	}
	r = make(map[[26]byte][]string)

	return result
}
