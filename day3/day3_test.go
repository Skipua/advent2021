package day3

import (
	"io/ioutil"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	got := PowerConsumption(input)

	expected := 198
	if got != expected {
		t.Errorf("Expected %v, got: %v", expected, got)
	}
}

func TestPowerConsumptionFromInput(t *testing.T) {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		t.Errorf("Couldn't read contents of input.txt")
	}

	input := strings.Split(strings.TrimSpace(string(file)), "\n")

	got := PowerConsumption(input)
	t.Logf("Power Consumption: %v", got) // 2724524 part 1

	if got != 2724524 {
		t.Errorf("Part1 Wrong! got: %v, should: %v", got, 2724524)
	}
}
