package day13

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	testInput(t, "example.txt", 16)
}

func TestInput(t *testing.T) {
	testInput(t, "input.txt", 17)
}

func testInput(t *testing.T, filename string, want int) {
	open, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("Couldn't read contents of input.txt")
	}

	rawInput := strings.Split(strings.TrimSpace(string(open)), "\n\n")
	dotsRaw := rawInput[0]
	foldsRaw := rawInput[1]

	got := CountFolds(strings.Split(dotsRaw, "\n"), strings.Split(foldsRaw, "\n"))

	t.Logf("Result: %v", got)

	if got != want {
		t.Errorf("Wrong! got: %v, should: %v", got, want)
	}
}
