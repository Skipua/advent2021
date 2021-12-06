package day5

import "math"

type Hydro struct {
	coveredPoints map[Point]int
}

type Point struct {
	x int
	y int
}

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func max(x1 int, x2 int) int {
	return int(math.Max(float64(x1), float64(x2)))
}

func min(x1 int, x2 int) int {
	return int(math.Min(float64(x1), float64(x2)))
}

func NewLine(x1 int, y1 int, x2 int, y2 int) Line {
	return Line{x1: x1, y1: y1, x2: x2, y2: y2}
}

func NewHydro() *Hydro {
	return &Hydro{make(map[Point]int)}
}

func (h *Hydro) AddLine(x1, y1, x2, y2 int) {
	line := NewLine(x1, y1, x2, y2)

	for _, p := range line.points() {
		if _, ok := h.coveredPoints[p]; ok {
			h.coveredPoints[p]++
		} else {
			h.coveredPoints[p] = 1
		}
	}
}

func (l *Line) points() []Point {
	points := make([]Point, 0)
	if l.x1 == l.x2 {
		for y := min(l.y1, l.y2); y <= max(l.y1, l.y2); y++ {
			points = append(points, Point{l.x1, y})
		}
	} else if l.y1 == l.y2 {
		for x := min(l.x1, l.x2); x <= max(l.x1, l.x2); x++ {
			points = append(points, Point{x, l.y1})
		}
	} else {
		var start Point
		var end Point
		if l.x1 < l.x2 {
			start = Point{l.x1, l.y1}
			end = Point{l.x2, l.y2}
		} else {
			start = Point{l.x2, l.y2}
			end = Point{l.x1, l.y1}
		}

		if end.y > start.y {
			for start != end {
				points = append(points, start)
				start = Point{start.x + 1, start.y + 1}
			}
		} else {
			for start != end {
				points = append(points, start)
				start = Point{start.x + 1, start.y - 1}
			}
		}

		points = append(points, end)
	}
	return points
}

func (h *Hydro) Result() interface{} {
	result := 0
	for _, v := range h.coveredPoints {
		if v >= 2 {
			result++
		}
	}
	return result
}
