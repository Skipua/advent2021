package day2

import (
	"strconv"
	"strings"
)

type Submarine struct {
	pos   int
	depth int
}

func NewSubmarine() *Submarine {
	return &Submarine{0, 0}
}

func (s *Submarine) Steer(command string) {
	commandSplitted := strings.Split(command, " ")
	direction := commandSplitted[0]
	x, _ := strconv.Atoi(commandSplitted[1])
	switch direction {
	case "forward":
		s.pos += x
	case "down":
		s.depth += x
	case "up":
		s.depth -= x
	}
}

func (s Submarine) Result() int {
	return s.pos * s.depth
}
