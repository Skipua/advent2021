package day9

import (
	"sort"
)

type Location struct {
	i int
	j int
}

func MeasureRisk(floor [][]int) int {
	rowLen := len(floor[0])

	isLowPoint := func(l Location) bool {
		i := l.i
		j := l.j
		h := floor[i][j]

		return (j == 0 || h < floor[i][j-1]) &&
			(j == rowLen-1 || h < floor[i][j+1]) &&
			(i == 0 || h < floor[i-1][j]) &&
			(i == len(floor)-1 || h < floor[i+1][j])
	}
	basins := make([]int, 0)
	for i, row := range floor {
		for j, _ := range row {
			location := Location{i, j}
			if isLowPoint(location) {
				visited := make(map[Location]struct{}, 0)
				visitBasin(floor, visited, location)
				basins = append(basins, len(visited))
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	return basins[0] * basins[1] * basins[2]
}

func visitBasin(floor [][]int, visited map[Location]struct{}, l Location) {
	if l.i < 0 || l.j < 0 || l.i == len(floor) || l.j == len(floor[0]) || floor[l.i][l.j] == 9 {
		return
	}
	visited[l] = struct{}{}
	for _, next := range []Location{
		{l.i + 1, l.j},
		{l.i - 1, l.j},
		{l.i, l.j - 1},
		{l.i, l.j + 1}} {
		if _, ok := visited[next]; !ok {
			visitBasin(floor, visited, next)
		}
	}
}
