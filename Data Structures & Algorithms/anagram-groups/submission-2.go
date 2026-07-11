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
		r[k] = append(r[k], str)
	}
	result := make([][]string, len(r))
	i := 0
	for _, v := range r {
		result[i] = v
		i++
	}
	return result
}
