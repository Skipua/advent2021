package day1

const minCountOfDigits = 4

func MeasureIncrements(in []int) int {
	if len(in) < minCountOfDigits {
		return 0
	}

	count := 0
	for i := 3; i < len(in); i++ {
		if in[i-3] < in[i] {
			count++
		}
	}

	return count
}
