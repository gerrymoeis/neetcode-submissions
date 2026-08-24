func heightChecker(heights []int) int {
    expected := make([]int, len(heights))
	copy(expected, heights[:])
	sort.Ints(expected)

	wrongs := 0
	for i, height := range heights {
		if height != expected[i] {
			wrongs++
		}
	}
	return wrongs
}