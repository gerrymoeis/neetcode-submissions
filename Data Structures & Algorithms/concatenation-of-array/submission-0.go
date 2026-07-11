func getConcatenation(nums []int) []int {
    n := len(nums)
	r := make([]int, n*2)
	for i, num := range nums {
		r[i] = num
		r[i+n] = num
	}
	return r
}
