package day10

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	testInput(t, "example.txt", 288957)
}

func TestFromInput(t *testing.T) {
	testInput(t, "input.txt", 2377613374)
}

func testInput(t *testing.T, filename string, want int) {
	open, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("Couldn't read contents of input.txt")
	}

	rawInput := strings.Split(strings.TrimSpace(string(open)), "\n")
	got := CheckSyntax(rawInput)
	t.Logf("Result: %v", got)

	if got != want {
		t.Errorf("Wrong! got: %v, should: %v", got, want)
	}
}
