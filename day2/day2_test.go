package day2

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestSubmarine_Steer(t *testing.T) {
	type args struct {
		commands []string
	}
	tests := []struct {
		name string
		args args
		res  int
	}{
		{
			"example_from_site",
			args{[]string{
				"forward 5",
				"down 5",
				"forward 8",
				"up 3",
				"down 8",
				"forward 2"}},
			900,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSubmarine()

			for _, command := range tt.args.commands {
				s.Steer(command)
			}

			got := s.Result()
			if got != tt.res {
				t.Errorf("Expected %v, got: %v", tt.res, got)
			}
		})
	}
}

func TestInput(t *testing.T) {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		t.Errorf("Couldn't read contents of input.txt")
	}

	input := strings.Split(strings.TrimSpace(string(file)), "\n")
	submarine := NewSubmarine()
	for _, command := range input {
		submarine.Steer(command)
	}

	got := submarine.Result()

	t.Logf("Got: %v", got)
}
