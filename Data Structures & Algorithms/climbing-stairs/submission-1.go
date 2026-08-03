func climbStairs(n int) int {
	first := 0
	second := 1
	for i := 0; i < n; i++ {
		temp := first + second
		first = second
		second = temp
	}
	return second
}