package day4

type Bingo struct {
	boards []Board
}

type Board struct {
	matrix  [][]int
	visited [][]bool
}

func NewBingo() *Bingo {
	boards := make([]Board, 0)
	return &Bingo{boards}
}

func NewBoard(size int) *Board {
	matrix := make([][]int, size)
	visited := make([][]bool, size)
	for i, _ := range matrix {
		matrix[i] = make([]int, size)
		visited[i] = make([]bool, size)
	}
	return &Board{matrix, visited}
}

func (board *Board) PopulateRow(row int, values []int) {
	for i, v := range values {
		board.matrix[row][i] = v
	}
}

func (bingo *Bingo) AddBoard(b Board) {
	bingo.boards = append(bingo.boards, b)
}

func (bingo *Bingo) Result() (int, bool) {
	return 0, false
}

func (board *Board) Mark(number int) (int, int, bool) {
	for rowIdx, row := range board.matrix {
		for colIdx, v := range row {
			if v == number {
				board.visited[rowIdx][colIdx] = true
				return rowIdx, colIdx, true
			}
		}
	}
	return 0, 0, false
}

func (board *Board) IsBingo(row int, col int) (Board, bool) {
	var noBoard Board
	allRowVisited := true
	for _, visited := range board.visited[row] {
		if !visited {
			allRowVisited = false
			break
		}
	}

	if allRowVisited {
		return *board, true
	}

	allColVisited := true
	for i := 0; i < len(board.visited); i++ {
		if !board.visited[i][col] {
			allColVisited = false
			break
		}
	}

	if allColVisited {
		return *board, true
	}

	return noBoard, false
}

func (board *Board) SumNotVisited() int {
	sum := 0
	for r, row := range board.visited {
		for c, v := range row {
			if !v {
				sum += board.matrix[r][c]
			}
		}
	}
	return sum
}

func (bingo *Bingo) NextNumber(number int) (Board, bool) {
	var winningBoard Board
	for _, b := range bingo.boards {
		row, col, marked := b.Mark(number)
		if !marked {
			continue
		}
		if winningBoard, isBingo := b.IsBingo(row, col); isBingo {
			return winningBoard, true
		}

	}
	return winningBoard, false
}
