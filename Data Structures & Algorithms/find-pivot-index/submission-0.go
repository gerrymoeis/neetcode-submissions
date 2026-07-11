func pivotIndex(nums []int) int {
	var t, tL, l int
	for _, n := range nums {
		t += n
	}
	for _, n := range nums {
		if tL == t-tL-n {
			return l
		}
		tL += n
		l++
	}
	return -1
}
