package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestNewSudokuBoard(t *testing.T) {
	n := 4
	s := New(n)
	if !(sort.IntsAreSorted(s.digits) && s.digits[n-1] == n) {
		t.Errorf("digits slice should be sorted and last digit should be %d", s.n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s.board[i][j] != 0 {
				t.Errorf("expected 0 got %v", s.board[i][j])
			}
		}
	}
}

func TestFixRow(t *testing.T) {
	n := 4
	s := New(n)
	s.board[2] = []int{3, 3, 2, 2}
	row := s.board[2]
	s.fixRow(2, false)
	fixedRow1 := []int{3, 1, 2, 4}
	fixedRow2 := []int{3, 4, 2, 1}
	if !(reflect.DeepEqual(row, fixedRow1) || reflect.DeepEqual(row, fixedRow2)) {
		t.Errorf("did not replace all twos: %v", row)
	}
	if !s.modified {
		t.Errorf("expected %t but got %t", true, false)
	}
}

func TestFixCol(t *testing.T) {
	n := 4
	s := New(n)
	s.board[0] = []int{1, 2, 1, 4}
	s.board[1] = []int{2, 3, 1, 4}
	s.board[2] = []int{3, 2, 3, 3}
	s.board[3] = []int{4, 3, 3, 1}

	colIndex := 2
	s.fixCol(colIndex, false)

	fixedCol1 := []int{2, 1, 4, 3}
	fixedCol2 := []int{4, 1, 2, 3}

	col := make([]int, s.n)
	for i := 0; i < s.n; i++ {
		col[i] = s.board[i][colIndex]
	}

	if !(reflect.DeepEqual(col, fixedCol1) || reflect.DeepEqual(col, fixedCol2)) {
		t.Errorf("did not replace all twos: %v", col)
	}

	t.Logf("%v", col)

	if !s.modified {
		t.Errorf("expected %t but got %t", true, false)
	}
}
