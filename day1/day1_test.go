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
		{"two", args{[]int{1, 1}}, 0},
		{"one increment", args{[]int{1, 2}}, 1},
		{"two increments", args{[]int{1, 2, 3}}, 2},
		{"two increments", args{[]int{1, 2, 3, 2, 3}}, 3},
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

	inputStrings := strings.Split(string(file), "\n")
	inputStrings = inputStrings[:len(inputStrings)-1]
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
