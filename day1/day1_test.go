package day1

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestMeasureIncrements(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{[]int{}}, 0},
		{"single", args{[]int{1}}, 0},
		{"two", args{[]int{1, 2}}, 0},
		{"three", args{[]int{1, 2, 3}}, 0},
		{"fourd groups", args{[]int{1, 2, 3, 4}}, 1}, // 1 + 2 + 3 < 2 + 3 + 4
		{"sliding", args{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MeasureIncrements(tt.args.in); got != tt.want {
				t.Errorf("MeasureIncrements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMeasureIncrementsFromInput(t *testing.T) {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		t.Errorf("Couldn't read contents of input.txt")
	}

	inputStrings := strings.Split(strings.TrimSpace(string(file)), "\n")
	input := make([]int, 0)
	for _, v := range inputStrings {
		digit, err := strconv.Atoi(v)
		if err != nil {
			t.Errorf("Couldn't convert to digit v=%v", v)
		}
		input = append(input, digit)
	}

	got := MeasureIncrements(input)
	t.Logf("Got increments: %v", got)
}
