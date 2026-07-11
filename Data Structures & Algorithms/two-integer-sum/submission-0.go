func twoSum(nums []int, target int) [2]int {
    h := make(map[int][2]int)
	for i := 0; i < len(nums); i++ {
		if nums[i] + h[nums[i]][0] == target && h[nums[i]][1] != i {
			return [2]int{h[nums[i]][1], i}
		}
		h[target - nums[i]] = [2]int{nums[i], i}
	}
	return [2]int{0,0}
}
