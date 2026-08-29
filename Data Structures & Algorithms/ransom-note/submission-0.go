func canConstruct(ransomNote string, magazine string) bool {
	rMap := make(map[byte]int)
	mMap := make(map[byte]int)
	for i := range ransomNote {
		rMap[ransomNote[i]]++
	}
	for i := range magazine {
		mMap[magazine[i]]++
	}
	for i := range ransomNote {
		if mMap[ransomNote[i]] < rMap[ransomNote[i]] {
			return false
		}
	}
	return true
}
