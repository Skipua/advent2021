package day4

import (
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

	input := strings.Split(strings.TrimSpace(string(open)), "\n\n")
	drawnNumbers := strings.Split(input[0], ",")

	bingo := NewBingo()
	for _, boardInput := range input[1:] {
		board := NewBoard(5)

		for rowIdx, row := range strings.Split(boardInput, "\n") {
			numbersInRow := make([]int, 5)
			for i, rowNum := range strings.Fields(row) {
				num, _ := strconv.Atoi(rowNum)
				numbersInRow[i] = num
			}
			board.PopulateRow(rowIdx, numbersInRow)
		}

		bingo.AddBoard(board)
	}

	var got int

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
