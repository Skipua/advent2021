package day5

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name   string
		args   [][]int
		result int
	}{
		{
			"intersects #1",
			[][]int{
				{0, 5, 10, 5},
				{2, 0, 2, 10},
			},
			1,
		},
		{
			"intersects #2",
			[][]int{
				{2, 0, 2, 10},
				{0, 5, 10, 5},
			},
			1,
		},
		{
			"all vertical overlap",
			[][]int{
				{0, 5, 0, 10},
				{0, 5, 0, 10},
			},
			6,
		},
		{
			"all horizontal overlap",
			[][]int{
				{5, 0, 10, 0},
				{5, 0, 10, 0},
			},
			6,
		},

		{
			"all horizontal overlap",
			[][]int{
				{2, 0, 4, 2},
				{0, 2, 2, 0},
			},
			1,
		},
		{
			"from reddit",
			[][]int{
				{70, 100, 120, 100},
				{70, 100, 500, 100},
			},
			51,
		},
		{
			"line overlaps but segment do not",
			[][]int{
				{0, 3, 3, 0},
				{0, 0, 2, 2},
			},
			0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := NewHydro()
			for _, points := range test.args {
				h.AddLine(points[0], points[1], points[2], points[3])
			}
			got := h.Result()
			if got != test.result {
				t.Errorf("Wanted: %v but got: %v", test.result, got)
			}
		})
	}
}

func TestFewPointsOverlap(t *testing.T) {

	hydro := NewHydro()
	hydro.AddLine(0, 9, 5, 9)
	hydro.AddLine(0, 9, 2, 9)
	result := hydro.Result()

	if result != 3 {
		t.Errorf("Got: %v, Should: %v", result, 3)
	}
}

func TestShouldBeNoOverlap(t *testing.T) {
	funcName(t)
}

func funcName(t *testing.T) {
	hydro := NewHydro()
	hydro.AddLine(0, 9, 5, 9)
	hydro.AddLine(6, 9, 10, 9)
	result := hydro.Result()

	want := 0
	if result != want {
		t.Errorf("Got: %v, Should: %v", result, want)
	}
}

func TestOneIntersection(t *testing.T) {

	hydro := NewHydro()
	hydro.AddLine(0, 0, 0, 5)
	hydro.AddLine(0, 0, 10, 0)
	result := hydro.Result()

	if result != 1 {
		t.Errorf("Got: %v, Should: %v", result, 1)
	}
}

func TestSinglePointOverlap(t *testing.T) {

	hydro := NewHydro()
	hydro.AddLine(5, 9, 10, 9)
	hydro.AddLine(0, 9, 5, 9)
	result := hydro.Result()

	if result != 1 {
		t.Errorf("Got: %v, Should: %v", result, 1)
	}
}

func TestFromInput(t *testing.T) {
	open, err := ioutil.ReadFile("input.txt")
	if err != nil {
		t.Errorf("Couldn't read contents of input.txt")
	}

	input := strings.Split(strings.TrimSpace(string(open)), "\n")
	h := NewHydro()

	for _, v := range input {
		lineInput := strings.Split(v, " -> ")
		lineStart := strings.Split(lineInput[0], ",")
		lineEnd := strings.Split(lineInput[1], ",")
		x1, _ := strconv.Atoi(lineStart[0])
		y1, _ := strconv.Atoi(lineStart[1])
		x2, _ := strconv.Atoi(lineEnd[0])
		y2, _ := strconv.Atoi(lineEnd[1])
		h.AddLine(x1, y1, x2, y2)
	}

	got := h.Result()

	t.Logf("Result: %v", got)

	want := 0
	if got != want {
		t.Errorf("Wrong! got: %v, should: %v", got, want)
	}
}
