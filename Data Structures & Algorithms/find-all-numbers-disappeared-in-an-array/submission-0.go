func findDisappearedNumbers(nums []int) []int {
	numMap := make(map[int]bool)
	for _, num := range nums {
		numMap[num] = true
	}
	res := []int{}
	for i := 1; i <= len(nums); i++ {
		if _, ok := numMap[i]; !ok {
			res = append(res, i)
		}
	}
	return res
}
