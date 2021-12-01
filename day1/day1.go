package day1

func MeasureIncrements(in []int) int {
	if len(in) < 2 {
		return 0
	}

	curr := in[0]
	count := 0
	for _, v := range in {
		if v > curr {
			count++
		}
		curr = v
	}

	return count
}
