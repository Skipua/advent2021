package day7

import "math"

func AlignCrabs(crabPositions []int) int {
	minFuel := math.MaxInt
	for _, c1 := range crabPositions {
		fuel := 0
		for _, c2 := range crabPositions {
			fuel += int(math.Abs(float64(c1) - float64(c2)))
		}

		if fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel
}
