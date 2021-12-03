package day3

import (
	"math"
	"strconv"
)

func PowerConsumption(input []string) int {
	total := len(input)
	bitSize := len(input[0])
	gammaStr := ""

	for i := 0; i < bitSize; i++ {
		num := 0
		for _, v := range input {
			bit, _ := strconv.Atoi(string(v[i]))
			num += bit
		}

		if num >= total/2 {
			gammaStr += "1"
		} else {
			gammaStr += "0"
		}
	}

	gamma := toInt(gammaStr)
	epsilon := gamma ^ (int(math.Pow(2, float64(bitSize))) - 1)
	return gamma * epsilon
}

func toInt(strBits string) int {
	parseInt, _ := strconv.ParseInt(strBits, 2, 32)
	return int(parseInt)
}
