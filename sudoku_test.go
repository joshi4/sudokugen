package main

import (
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
