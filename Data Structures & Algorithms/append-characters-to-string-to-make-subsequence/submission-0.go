func appendCharacters(s string, t string) int {
    i := 0
	for _, c := range s {
		if i >= len(t) {
			break
		}
		if t[i] == byte(c) {
			i++
		}
	}
	return len(t) - i
}