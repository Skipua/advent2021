package day6

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
		days  int
		want  int
	}{
		{
			"simple",
			[]int{3},
			4,
			2,
		},
		{
			"simple #",
			[]int{3},
			13,
			4,
		},
		{
			"simple",
			[]int{3, 4},
			5,
			4,
		},
		{
			"simple",
			[]int{3, 4, 3, 1, 2},
			18,
			26,
		},
		{
			"simple",
			[]int{3, 4, 3, 1, 2},
			80,
			5934,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := SimulateLantern(test.input, test.days)
			if got != test.want {
				t.Errorf("Want: %v; Got: %v", test.want, got)
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

	got := SimulateLantern(input, 256)
	t.Logf("Result: %v", got)

	want := 1622533344325
	if got != want {
		t.Errorf("Wrong! got: %v, should: %v", got, want)
	}
}
