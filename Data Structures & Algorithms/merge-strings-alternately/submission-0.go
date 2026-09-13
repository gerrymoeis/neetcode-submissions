import "unsafe"

func mergeAlternately(word1 string, word2 string) string {
	m := make([]byte, len(word1) + len(word2))
	var i, j int
	for j < len(word1) || j < len(word2) {
		if j < len(word1) {
			m[i] = word1[j]
			i++
		}
		if j < len(word2) {
			m[i] = word2[j]
			i++
		}
		j++
	}
	return unsafe.String(&m[0], len(m))
}
