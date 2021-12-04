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
			bingo.AddBoard(board)
		}
	}

	got := 0
	for _, v := range drawnNumbers {
		num, _ := strconv.Atoi(v)

		bingo.NextNumber(num)

		if bingo.WonBoardsCount() == bingo.BoardsCount() {
			lastWinningBoard := bingo.lastWonBoard
			t.Logf("Last winning board: %v", lastWinningBoard)
			got = lastWinningBoard.SumNotVisited() * num
			break
		}
	}

	t.Logf("Result: %v", got)

	//part 1
	//want := 38913
	want := 16836
	if got != want {
		t.Errorf("Wrong! got: %v, should: %v", got, want)
	}
}
