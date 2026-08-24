func maxNumberOfBalloons(text string) int {
    b := make([]int, 7)
	var countL, countO bool
	for _, c := range text {
		switch c {
			case 'b':
				b[0]++
			case 'a':
				b[1]++
			case 'l':
				if countL == true {
					b[3]++
					countL = false
				} else {
					b[2]++
					countL = true
				}
			case 'o':
				if countO == true {
					b[4]++
					countO = false
				} else {
					b[5]++
					countO = true
				}
			case 'n':
				b[6]++
		}
	}
	min := math.MaxInt
	for _, n := range b {
		if n < min {
			min = n
		}
	}
	return min
}