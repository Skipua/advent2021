package day7

import "math"

func AlignCrabs(crabPositions []int) int {
	minFuel := math.MaxInt
	for _, initPos := range crabPositions {
		fuel := 0
		for _, targetPos := range crabPositions {
			fuel += calculateFuel(initPos, targetPos)
		}

		if fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel
}

func calculateFuel(pos1 int, pos2 int) int {
	fuel := 0
	diffBetweenPositions := int(math.Abs(float64(pos1) - float64(pos2)))
	if diffBetweenPositions == 0 {
		return 0
	}

	for i := 1; i <= diffBetweenPositions; i++ {
		fuel += i
	}
	return fuel
}
