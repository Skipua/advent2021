package day3

import (
	"strconv"
)

func PowerConsumption(input []string) int {
	oxy := findMostOrLeastCommon(0, input, true)
	co2 := findMostOrLeastCommon(0, input, false)
	return toInt(oxy) * toInt(co2)
}

func findMostOrLeastCommon(idx int, input []string, mostCommon bool) string {
	if len(input) == 1 {
		return input[0]
	} else {
		mostCommonBit := mostCommonBitAt(input, idx)
		newInput := make([]string, 0)
		for _, v := range input {
			if mostCommon && string(v[idx]) == strconv.Itoa(mostCommonBit) {
				newInput = append(newInput, v)
			} else if !mostCommon && string(v[idx]) == strconv.Itoa(1-mostCommonBit) {
				newInput = append(newInput, v)
			}
		}
		idx++
		return findMostOrLeastCommon(idx, newInput, mostCommon)
	}
}

func mostCommonBitAt(input []string, idx int) int {
	num := 0
	for _, v := range input {
		bit, _ := strconv.Atoi(string(v[idx]))
		num += bit
	}

	half := float64(len(input)) / 2.0
	if float64(num) >= half {
		return 1
	} else {
		return 0
	}
}

func toInt(strBits string) int {
	parseInt, _ := strconv.ParseInt(strBits, 2, 32)
	return int(parseInt)
}
