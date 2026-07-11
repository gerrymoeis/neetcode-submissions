type NumArray struct {
    nums []int
}


func Constructor(numsArr []int) NumArray {
    return NumArray{
		nums: numsArr,
	}
}


func (this *NumArray) SumRange(left int, right int) int {
    total := 0
	for _, num := range this.nums[left:right+1] {
		total += num
	}
	return total
}