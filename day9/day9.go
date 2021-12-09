package day9

func MeasureRisk(floor [][]int) int {
	rowLen := len(floor[0])

	isLowPoint := func(i, j int) bool {
		h := floor[i][j]

		return (j == 0 || h < floor[i][j-1]) &&
			(j == rowLen-1 || h < floor[i][j+1]) &&
			(i == 0 || h < floor[i-1][j]) &&
			(i == len(floor)-1 || h < floor[i+1][j])
	}
	risk := 0
	for i, row := range floor {
		for j, height := range row {
			if isLowPoint(i, j) {
				risk += height + 1
			}
		}
	}
	return risk
}
