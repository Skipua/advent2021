package day13

import (
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

type Paper struct {
	paper []uint16
	bits  int
}

type Dot struct {
	x int
	y int
}

func NewPaper(n, bits int) *Paper {
	paper := make([]uint16, n)
	for i := 0; i < len(paper); i++ {
		paper[i] = 0
	}
	return &Paper{paper, bits}
}

func (p *Paper) addDots(dots []Dot) {
	for _, dot := range dots {
		p.paper[dot.y] |= 1 << uint16(dot.x)
	}
}

func (p *Paper) foldUp(y int) {
	for i := 0; i < y; i++ {
		p.paper[i] |= p.paper[len(p.paper)-1-i]
	}
	p.paper = p.paper[:y]
}

func (p *Paper) foldLeft(x int) {
	for i, v := range p.paper {
		reversed := bits.Reverse16(v)
		toShift := 16 - p.bits
		mirrored := reversed >> toShift
		p.paper[i] = (p.paper[i] | mirrored) >> (x + 1)
	}
	p.bits -= x + 1
}

func (p Paper) String() string {
	str := ""

	for _, num := range p.paper {
		formatUint := strconv.FormatUint(uint64(num), 2)
		str += strings.Repeat("0", p.bits-len(formatUint)) + formatUint + "\n"
	}

	return str
}

func (p Paper) countDots() int {
	count := 0

	for _, v := range p.paper {
		count += bits.OnesCount16(v)
	}

	return count
}

func CountFolds(rawDots, rawFolds []string) int {
	dots := toListOfDots(rawDots)
	paper := NewPaper(maxY(dots)+1, maxX(dots)+1)
	paper.addDots(dots)

	for i, f := range rawFolds {
		xOrY, index := parseFold(f)

		fmt.Println(paper)
		if xOrY == "y" {
			paper.foldUp(index)
		} else {
			paper.foldLeft(index)
		}

		fmt.Printf("AFTER %v=%v %v\n", xOrY, index, i+1)
		fmt.Printf("COUNT AFTER %v: %v\n", i+1, paper.countDots())

		fmt.Println(paper)
	}

	return paper.countDots()
}

func maxX(dots []Dot) int {
	max := dots[0].x

	for i := 1; i < len(dots); i++ {
		if max < dots[i].x {
			max = dots[i].x
		}
	}

	return max
}

func maxY(dots []Dot) int {
	max := dots[0].y

	for i := 1; i < len(dots); i++ {
		if max < dots[i].y {
			max = dots[i].y
		}
	}

	return max
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
