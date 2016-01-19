package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

type Sudoku struct {
	n        int
	modified bool
	board    [][]int
	digits   []int
}

func init() {
	now := time.Now()
	rand.Seed(now.Unix())
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
	// defer wg.Done()
	// permute the indices ( )
	// so don't need to add 1 every time
	permuteIndices := rand.Perm(s.n)

	N := int(math.Sqrt(float64(s.n)))
	rowIndex := (squareIndex / N) * N
	colIndex := (squareIndex % N) * N
	rowSlice := s.board[rowIndex : rowIndex+N]

	for i := range rowSlice {
		for j := colIndex; j < colIndex+N; j++ {
			rowSlice[i][j] = s.digits[permuteIndices[N*i+(j%N)]]
		}
	}
}

// FixRow(i) // change all but first conflict
func (s *Sudoku) FixRows() {
	for i := 0; i < s.n; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			s.fixRow(index)
		}(i)
	}
	wg.Wait()
}

// FixColumn(i) //change all but last conflict

// if there is an error in a row i:e  number(s) is(are) repeated or
// conversely if number(s) is(are) missing this function fixes it.
// fixRow will change all but first conflict
//
// Example 1
// n :=  4 and a row has 1 2 2 4
// Then  the row will be changed to 1 2 3 4 ( keeping first 2 and replacing all others )
//
// Example 2:
// n := 4 and a row has 1 2 2 2
// row will be changed to 1 2 3 4 or 1 2 4 3 ( missing numbers are substitued randomly for extra ones )
//
// Note: fixColumn will change all but last conflict
// This convention is important to avoid being stuck indefinitely by changes that cancel each other out.
func (s *Sudoku) fixRow(rowIndex int) {
	// map of digit to all indices it is repeated in
	digitCount := make(map[int][]int)
	missingDigits := make([]int, 0)
	repeatedDigits := make(map[int]struct{})

	row := s.board[rowIndex]
	for i := range row {
		digit := row[i]
		if _, ok := digitCount[digit]; !ok {
			digitCount[digit] = make([]int, 0)
		} else if len(digitCount[digit]) >= 1 {
			repeatedDigits[digit] = struct{}{}
		}
		digitCount[digit] = append(digitCount[digit], i)
	}

	for _, digit := range s.digits {
		if _, ok := digitCount[digit]; !ok {
			missingDigits = append(missingDigits, digit)
		}
	}

	randomIndexes := rand.Perm(len(missingDigits))
	index := 0
	for repeatedDigit, _ := range repeatedDigits {
		for _, indexToReplace := range digitCount[repeatedDigit][1:] {
			row[indexToReplace] = missingDigits[randomIndexes[index]]
			index += 1

			if !s.modified {
				s.modified = true
			}
		}
	}
}

func (s *Sudoku) DisplayBoard() {
	for i := range s.board {
		fmt.Printf("%v\n", s.board[i])
	}
}

var wg sync.WaitGroup

func main() {
	s := New(4)
	for i := 0; i < s.n; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			s.PopulateSquare(index)
		}(i)
	}
	wg.Wait()

	s.DisplayBoard()
	s.FixRows()
	fmt.Println()
	s.DisplayBoard()
}
