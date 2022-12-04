package util

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(slice []int) (min int) {
	if len(slice) == 0 {
		return
	}

	min = slice[0]
	for _, v := range slice {
		if v < min {
			min = v
		}
	}

	return
}

func Max(slice []int) (max int) {
	if len(slice) == 0 {
		return
	}

	max = slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}

	return
}
