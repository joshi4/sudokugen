package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Sudoku struct {
	n        int
	modified bool
	board    [][]int
	digits   []int
}

func New(n int) *Sudoku {
	digits := make([]int, n, n)
	for i := 1; i <= n; i++ {
		digits[i-1] = i
	}

	board := make([][]int, n, n)
	for i := range board {
		board[i] = make([]int, n, n)
	}
	return &Sudoku{
		n:        n,
		modified: false,
		board:    board,
		digits:   digits,
	}
}

// Given a squareIndex: initialize relevant grid in s.board
// to a random permuation of 1.. s.n
// index = 0: rows: 0 2 cols: 0 -2
// index = 1: rows: 0 -2 cols 3 - 5
// index = 2: rows 0 - 2 cols 6 - 8
// index = 3: rows 3 - 5: cols 0 - 2

// Formula: iNdex = k: rows (k/N)*N : (add N to prev result) cols: (k %N) *N : (add N to previous result)
// where N = sqrt(n)
func (s *Sudoku) PopulateSquare(squareIndex int) {
	// permute the indices ( )
	// so don't need to add 1 every time
	permute_indices := rand.Perm(s.n)
	N := int(math.Sqrt(float64(s.n)))
	row_index := (squareIndex / N) * N
	col_index := (squareIndex % N) * N
	row_slice := s.board[row_index : row_index+N]
	for i := range row_slice {
		for j := col_index; j < col_index+N; j++ {
			row_slice[i][j] = s.digits[permute_indices[N*i+(j%N)]]
		}
	}
}

func (s *Sudoku) DisplayBoard() {
	for i := range s.board {
		fmt.Printf("%v\n", s.board[i])
	}
}

func main() {
	s := New(9)
	for i := 0; i < s.n; i++ {
		s.PopulateSquare(i)
	}
	s.DisplayBoard()
}
