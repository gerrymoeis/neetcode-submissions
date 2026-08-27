func largestGoodInteger(num string) string {
	var res, substring string
	max := 0
	for _, n := range num {
		if (len(substring) >= 1 && byte(n) != substring[0]) || len(substring) == 3 {
			val, _ := strconv.Atoi(substring)
			if len(substring) == 3 {
				if val >= max {
					max = val
					res = substring
				}
			}
			substring = ""
		}
		substring += string(n)
	}
	val, _ := strconv.Atoi(substring)
	if len(substring) == 3 {
		if val >= max {
			max = val
			res = substring
		}
	}
	return res
}
