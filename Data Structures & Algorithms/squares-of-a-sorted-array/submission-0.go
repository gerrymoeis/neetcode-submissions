import "slices"

func sortedSquares(nums []int) []int {
	for i := range nums {
		nums[i] *= nums[i]
	}
	slices.Sort(nums)
	return nums
}