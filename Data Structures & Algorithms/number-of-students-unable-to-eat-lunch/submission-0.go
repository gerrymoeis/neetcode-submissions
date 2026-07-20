func countStudents(students []int, sandwiches []int) int {
	countStud := make(map[int]int)
	for _, s := range students {
		countStud[s]++
	}
	result := len(students)
	for _, s := range sandwiches {
		if countStud[s] > 0 {
			countStud[s]--
			result--
		} else {
			break
		}
	}
	return result
}