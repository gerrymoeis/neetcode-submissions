func countSeniors(details []string) int {
    count := 0
	for _, d := range details {
		a := int(d[11])
		b := int(d[12])
		if a > 54 {
			count++
		} else if a == 54 && b > 48 {
			count++
		}
	}
	return count
}