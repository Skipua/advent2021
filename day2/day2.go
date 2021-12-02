package day2

import (
	"strconv"
	"strings"
)

type Submarine struct {
	hPos  int
	depth int
	aim   int
}

func NewSubmarine() *Submarine {
	return &Submarine{0, 0, 0}
}

func (s *Submarine) Steer(command string) {
	commandSplitted := strings.Split(command, " ")
	direction := commandSplitted[0]
	x, _ := strconv.Atoi(commandSplitted[1])
	switch direction {
	case "forward":
		s.hPos += x
		s.depth += s.aim * x
	case "up":
		s.aim -= x
	case "down":
		s.aim += x
	}
}

func (s Submarine) Result() int {
	return s.hPos * s.depth
}
