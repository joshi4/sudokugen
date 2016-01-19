package main

type Sudoku struct {
	n        int
	modified bool
	board    [][]int
}

func New(n int) *Sudoku {
	board := make([][]int, n, n)
	for i := range board {
		board[i] = make([]int, n, n)
	}
	return &Sudoku{
		n:        n,
		modified: false,
		board:    board,
	}
}

// Given a squareIndex: initialize relevant grid in s.board
// to a random permuation of 1.. s.n
func (s *Sudoku) PopulateSquare(squareIndex int) {

}
