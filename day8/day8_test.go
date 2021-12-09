package day8

import (
	"io/ioutil"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	doTestFromInputFIle(t, "testInput.txt", 5353)
}

func Test2(t *testing.T) {
	doTestFromInputFIle(t, "testInput2.txt", 9600)
}

func TestFromInput(t *testing.T) {
	doTestFromInputFIle(t, "input.txt", 0)
}

func doTestFromInputFIle(t *testing.T, fileName string, want int) {
	open, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Errorf("Couldn't read contents of input.txt")
	}

	entries := strings.Split(strings.TrimSpace(string(open)), "\n")
	display := NewDisplay()
	for _, entry := range entries {
		entryParsed := strings.Split(entry, " | ")
		signals := strings.Fields(entryParsed[0])
		fourDigitOutput := strings.Fields(entryParsed[1])

		display.Parse(signals, fourDigitOutput)
	}

	got := display.SumOfNumbers()
	t.Logf("Result: %v", got)

	if got != want {
		t.Errorf("Wrong! got: %v, should: %v", got, want)
	}
}
