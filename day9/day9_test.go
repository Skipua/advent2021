package day9

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	testInput(t, "example.txt", 15)
}

func TestFromInput(t *testing.T) {
	testInput(t, "input.txt", 0)
}

func testInput(t *testing.T, filename string, want int) {
	open, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("Couldn't read contents of input.txt")
	}

	rawInput := strings.Split(strings.TrimSpace(string(open)), "\n")
	input := make([][]int, len(rawInput))

	for i, raw := range rawInput {
		input[i] = make([]int, 0)
		for _, value := range strings.Split(raw, "") {
			height, _ := strconv.Atoi(value)
			input[i] = append(input[i], height)
		}
	}

	got := MeasureRisk(input)
	t.Logf("Result: %v", got)

	if got != want {
		t.Errorf("Wrong! got: %v, should: %v", got, want)
	}
}
