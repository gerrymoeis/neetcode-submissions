func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := range res {
		row := make([]int, i+1)
		for j := range row {
			row[j] = 1
		}
		for k := 1; k < i; k++ {
			row[k] = row[k-1] * (i-k+1) / k
		}
		res[i] = row
	}
	return res
}
