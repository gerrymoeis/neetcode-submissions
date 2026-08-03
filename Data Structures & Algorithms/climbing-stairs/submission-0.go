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

// func climbStairs(n int) int {
// 	if n == 0 {
// 		return 0
// 	}
// 	count := 1
// 	if n % 2 == 0 {
// 		count++
// 	}
// 	return count
// }

// func climbStairs(n int) int {
// 	count := 0
// 	for i := 0; i <= n; i++ {
// 		temp := count
// 		count -= temp
// 		count = temp + i
// 	}
// 	fmt.Println(count)
// 	return n
// }
