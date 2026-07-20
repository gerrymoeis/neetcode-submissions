func countStudents(students []int, sandwiches []int) int {
	countStud := make([]int, 2)
	for _, s := range students {
		countStud[s]++
	}
	for i, s := range sandwiches {
		if countStud[s] > 0 {
			countStud[s]--
		} else {
			return len(students) - i
		}
	}
	return 0
}