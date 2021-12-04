package day4

type Bingo struct {
	boards         []*Board
	wonBoardsCount int
	lastWonBoard   Board
}

type Board struct {
	matrix  [][]int
	visited [][]bool
	won     bool
}

func NewBingo() *Bingo {
	var empty Board
	boards := make([]*Board, 0)
	return &Bingo{boards, 0, empty}
}

func NewBoard(size int) *Board {
	matrix := make([][]int, size)
	visited := make([][]bool, size)
	for i, _ := range matrix {
		matrix[i] = make([]int, size)
		visited[i] = make([]bool, size)
	}
	return &Board{matrix, visited, false}
}

func (bingo *Bingo) WonBoardsCount() int {
	return bingo.wonBoardsCount
}

func (bingo *Bingo) BoardsCount() int {
	return len(bingo.boards)
}

func (board *Board) PopulateRow(row int, values []int) {
	copy(board.matrix[row], values)
}

func (bingo *Bingo) AddBoard(b *Board) {
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

func (board *Board) IsBingo(row int, col int) bool {
	return board.allRowVisited(row) || board.allColumnVisited(col)
}

func (board *Board) allColumnVisited(col int) bool {
	allColVisited := true
	for i := 0; i < len(board.visited); i++ {
		if !board.visited[i][col] {
			allColVisited = false
			break
		}
	}
	return allColVisited
}

func (board *Board) allRowVisited(row int) bool {
	allRowVisited := true
	for _, visited := range board.visited[row] {
		if !visited {
			allRowVisited = false
			break
		}
	}
	return allRowVisited
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

func (bingo *Bingo) NextNumber(number int) {
	for _, b := range bingo.boards {
		row, col, marked := b.Mark(number)
		if !marked || b.won {
			continue
		}
		if b.IsBingo(row, col) {
			b.won = true
			bingo.wonBoardsCount++
			bingo.lastWonBoard = *b
		}
	}
}
