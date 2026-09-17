func sortArrayByParity(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	var i, j int
	for _, num := range nums {
		if num%2 != 0 {
			res[n-j-1] = num
			j++
		} else {
			res[i] = num
			i++
		}
	}
	return res
}
