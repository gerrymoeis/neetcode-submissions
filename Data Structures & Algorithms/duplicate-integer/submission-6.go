func hasDuplicate(nums []int) bool {
    dups := make(map[int]bool)
        for _, num := range nums {
            if dups[num] {
                return true
            }
            dups[num] = true
        }
    return false
}