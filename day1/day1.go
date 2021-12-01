package day1

const minCountOfDigits = 4

func MeasureIncrements(in []int) int {
	if len(in) < minCountOfDigits {
		return 0
	}

	windowsCount := len(in) - 2
	count := 0
	prevSum := 0
	sum := 0
	for w := 0; w < windowsCount; w++ {
		for i := w; i < w+3; i++ {
			sum += in[i]
		}

		if prevSum > 0 && prevSum < sum {
			count++
		}

		prevSum = sum
		sum = 0
	}

	return count
}
