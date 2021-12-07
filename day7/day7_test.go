package day7

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		fuel  int
	}{
		{
			"example",
			[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			37,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotFuel := AlignCrabs(test.input)
			if gotFuel != test.fuel {
				t.Errorf("Wanted: %v, but got %v", test.fuel, gotFuel)
			}
		})
	}
}

func TestFromInput(t *testing.T) {
	open, err := ioutil.ReadFile("input.txt")
	if err != nil {
		t.Errorf("Couldn't read contents of input.txt")
	}

	rawInput := strings.Split(strings.TrimSpace(string(open)), ",")
	input := make([]int, len(rawInput))

	for i, raw := range rawInput {
		ttl, _ := strconv.Atoi(raw)
		input[i] = ttl
	}

	got := AlignCrabs(input)
	t.Logf("Result: %v", got)

	want := 1622533344325
	if got != want {
		t.Errorf("Wrong! got: %v, should: %v", got, want)
	}
}
