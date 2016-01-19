package main

import "testing"

func TestNewSudokuBoard(t *testing.T) {
	n := 4
	s := New(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s.board[i][j] != 0 {
				t.Errorf("expected 0 got %v", s.board[i][j])
			}
		}
	}
}
