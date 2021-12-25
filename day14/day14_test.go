package day14

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	//testInput(t, "example.txt", 1, 1)
	//testInput(t, "example.txt", 2, 5)
	//testInput(t, "example.txt", 3, 7)
	testInput(t, "example.txt", 4, 18)
	//testInput(t, "example.txt", 10, 1588)
}

func TestInputPart1(t *testing.T) {
	testInput(t, "input.txt", 10, 2408)
}

func TestInputPart2(t *testing.T) {
	testInput(t, "input.txt", 40, 0)
}

func testInput(t *testing.T, filename string, steps, want int) {
	open, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("Couldn't read contents of input.txt")
	}

	rawInput := strings.Split(strings.TrimSpace(string(open)), "\n\n")
	init := rawInput[0]
	polyInstructions := rawInput[1]

	got := DoPoly(init, strings.Split(polyInstructions, "\n"), steps)

	t.Logf("Result: %v", got)

	if got != want {
		t.Errorf("Wrong! got: %v, should: %v", got, want)
	}
}
