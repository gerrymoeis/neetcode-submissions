func replaceElements(arr []int) []int {
	n := len(arr)-1
	rightMax := arr[n]
	for i, _ := range arr[:n] {
		tempRM := arr[n-1-i]
		arr[n-1-i] = rightMax
		if tempRM > rightMax {
			rightMax = tempRM
		}
	}
	arr[n] = -1
	return arr
}
