package day11

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestFromExampleTenSteps(t *testing.T) {
	testInput(t, "example.txt", 10, 204)
}

func TestFromExample100Steps(t *testing.T) {
	testInput(t, "example.txt", 100, 1656)
}

func TestFromInput(t *testing.T) {
	testInput(t, "input.txt", 100, 1757)
}

func testInput(t *testing.T, filename string, steps, want int) {
	open, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("Couldn't read contents of input.txt")
	}

	rawInput := strings.Split(strings.TrimSpace(string(open)), "\n")
	input := make([][]int, len(rawInput))
	for i, raw := range rawInput {
		input[i] = make([]int, 0)
		for _, v := range strings.Split(raw, "") {
			num, _ := strconv.Atoi(v)
			input[i] = append(input[i], num)
		}
	}
	got := CountFlashes(input, steps)
	t.Logf("Result: %v", got)

	if got != want {
		t.Errorf("Wrong! got: %v, should: %v", got, want)
	}
}
