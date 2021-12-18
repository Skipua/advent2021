package day12

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestSmall(t *testing.T) {
	testInput(t, "small-example.txt", 36)
}

func TestFromInput(t *testing.T) {
	testInput(t, "input.txt", 92111)
}

func testInput(t *testing.T, filename string, want int) {
	open, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("Couldn't read contents of input.txt")
	}

	rawInput := strings.Split(strings.TrimSpace(string(open)), "\n")
	got := CountPaths(rawInput)
	t.Logf("Result: %v", got)

	if got != want {
		t.Errorf("Wrong! got: %v, should: %v", got, want)
	}
}
