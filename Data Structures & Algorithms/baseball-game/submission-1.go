func calPoints(operations []string) int {
	total := 0
	records := make(map[int]int)
	i := 0
processRecords:
	for _, op := range operations {
		switch op {
		case "+":
			records[i] = records[i-1] + records[i-2]
		case "D":
			records[i] = records[i-1] * 2
		case "C":
			i--
			total -= records[i]
			continue processRecords
		default:
			records[i], _ = strconv.Atoi(op)
		}
		total += records[i]
		i++
	}
	return total
}