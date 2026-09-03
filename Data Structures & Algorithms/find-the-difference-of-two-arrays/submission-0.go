func findDifference(nums1 []int, nums2 []int) [][]int {
	answer := make([][]int, 2)
	nums1Map := make(map[int]bool)
	nums2Map := make(map[int]bool)
	for _, num := range nums1 {
		nums1Map[num] = true
	}
	for _, num := range nums2 {
		nums2Map[num] = true
	}
	for _, num := range nums1 {
		if _, ok := nums2Map[num]; !ok {
			answer[0] = append(answer[0], num)
			nums2Map[num] = true
		}
	}
	for _, num := range nums2 {
		if _, ok := nums1Map[num]; !ok {
			answer[1] = append(answer[1], num)
			nums1Map[num] = true
		}
	}
	return answer
}