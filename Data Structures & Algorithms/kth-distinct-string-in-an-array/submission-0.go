func kthDistinct(arr []string, k int) string {
    arrMap := make(map[string]int)
	distinct := []string{}
	for _, c := range arr {
		arrMap[c]++
	}
	for _, c := range arr {
		if arrMap[c] == 1 {
			distinct = append(distinct, c)
		}
	}
	if len(distinct) < k {
		return ""
	}
	return distinct[k-1]
}