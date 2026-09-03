func signFunc(product int) int {
	switch {
	case product < 0:
		return -1
	case product > 0:
		return 1
	}
	return 0
}

func arraySign(nums []int) int {
	product := 1
	for _, num := range nums {
		if num == 0 {
			return signFunc(0)
		}
		sign := 1 | (num >> 63)
		product *= sign
	}
	return signFunc(product)
}
