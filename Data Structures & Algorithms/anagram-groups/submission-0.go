func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	l := [26]int{}
	for i := 0; i < len(s); i++ {
		l[s[i]-'a']++
		l[t[i]-'a']--
	}
	for _, v := range l {
		if v != 0 {
			return false
		}
	}
	return true
}

func groupAnagrams(strs []string) [][]string {
	r := make(map[string][]string)
	for _, str := range strs {
		a := false
		for k, _ := range r {
			if b := isAnagram(k, str); b {
				r[k] = append(r[k], str)
				a = b
			}
		}
		if !a {
			r[str] = []string{str}
		}
	}
	result := make([][]string, len(r))
	for k, v := range r {
		result[cap(result)-len(r)] = v
		delete(r, k)
	}
	return result
}
