func countCharacters(words []string, chars string) int {
    charMap := make(map[byte]int)
	for i := range chars {
		charMap[chars[i]]++
	}
	total := 0
	for _, word := range words {
		count := 0
		wordMap := make(map[byte]int)
		for i := range word {
			if n, ok := charMap[word[i]]; ok && n > 0 {
				count++
				if _, ok := wordMap[word[i]]; !ok {
					wordMap[word[i]] = n
				}
				charMap[word[i]]--
			} else {
				count = 0
				break
			}
		}
		for c, n := range wordMap {
			charMap[c] = n
		}
		total += count
	}
	return total
}