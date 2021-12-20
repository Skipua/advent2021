package day13

import (
	"fmt"
	"strconv"
	"strings"
)

const MAX_X = 895
const MAX_Y = 1311

type Paper struct {
	paper [][]bool
}

type Dot struct {
	x, y int
}

func NewPaper() *Paper {
	paper := make([][]bool, MAX_X)
	for i := 0; i < len(paper); i++ {
		paper[i] = make([]bool, MAX_Y)
		for j := 0; j < len(paper[i]); j++ {
			paper[i][j] = false
		}
	}
	return &Paper{paper}
}

func (p *Paper) addDots(dots []Dot) {
	for _, dot := range dots {
		p.paper[dot.y][dot.x] = true
	}
}

func (p *Paper) foldUp(y int) {
	for i := 0; i < y; i++ {
		p.paper[i] = overlap(p.paper[i], p.paper[len(p.paper)-1-i])
	}
	p.paper = p.paper[:y]
}

func (p *Paper) foldLeft(x int) {
	for i, row := range p.paper {
		p.paper[i] = overlap(row[:x], reverse(row)[:x])
	}
}

func reverse(v []bool) []bool {
	r := make([]bool, len(v))

	for i := 0; i < len(v); i++ {
		r[i] = v[len(v)-1-i]
	}

	return r
}

func overlap(a, b []bool) []bool {
	res := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = a[i] || b[i]
	}
	return res
}

func (p Paper) String() string {
	str := ""

	for _, row := range p.paper {
		for _, v := range row {
			if v {
				str += "⬜"
			} else {
				str += "⬛"
			}
		}
		str += "\n"
	}

	return str
}

func (p Paper) countDots() int {
	count := 0

	for _, row := range p.paper {
		for _, v := range row {
			if v {
				count++
			}
		}
	}

	return count
}

func CountFolds(rawDots, rawFolds []string) int {
	dots := toListOfDots(rawDots)

	paper := NewPaper()
	paper.addDots(dots)

	for _, f := range rawFolds {
		axis, index := parseFold(f)

		if axis == "y" {
			paper.foldUp(index)
		} else {
			paper.foldLeft(index)
		}
	}

	fmt.Println(paper)
	return paper.countDots()
}

func toListOfDots(rawDots []string) []Dot {
	dots := make([]Dot, len(rawDots))

	for i, v := range rawDots {
		split := strings.Split(v, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		dots[i] = Dot{x, y}
	}

	return dots
}

func parseFold(firstFold string) (string, int) {
	fold := strings.Split(strings.Fields(firstFold)[2], "=")
	num, _ := strconv.Atoi(fold[1])
	return fold[0], num
}
