package day4

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestFromInput(t *testing.T) {
	open, err := ioutil.ReadFile("input.txt")
	if err != nil {
		t.Errorf("Couldn't read contents of input.txt")
	}

	input := strings.Split(strings.TrimSpace(string(open)), "\n")
	drawnNumbers := strings.Split(input[0], ",")
	fmt.Println(drawnNumbers)

	bingo := NewBingo()
	board := NewBoard(5)
	row := 0
	for _, v := range input[2:] {
		if string(v) == "" {
			board = NewBoard(5)
			row = 0
			continue
		}

		rowStrings := strings.Fields(v)
		numbersInRow := make([]int, 5)
		for i, numStr := range rowStrings {
			num, _ := strconv.Atoi(numStr)
			numbersInRow[i] = num
		}
		board.PopulateRow(row, numbersInRow)
		row++
		if row == 5 {
			bingo.AddBoard(*board)
		}
	}

	got := 0
	for _, v := range drawnNumbers {
		num, _ := strconv.Atoi(v)

		board, isBingo := bingo.NextNumber(num)

		if isBingo {
			t.Logf("Bingo for board: %v", board)
			got = board.SumNotVisited() * num
			break
		}
	}

	t.Logf("We have a winning board with Result: %v", got)

	want := 38913
	if got != want {
		t.Errorf("Wrong! got: %v, should: %v", got, want)
	}
}
