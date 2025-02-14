package unit

import (
	"battleship/game"
	"testing"
)

func TestDisplayCell(t *testing.T) {
	board := game.Board{
		Grid: [10][10]rune{
			{'1', '0', '1'},
		},
		Hits: [10][10]bool{
			{false, true, false},
		},
	}

	tests := []struct {
		row, col     int
		expectedRune rune
	}{
		{0, 0, '1'},
		{0, 1, 'X'},
		{0, 2, '1'},
	}

	for _, tt := range tests {
		result := board.DisplayCell(tt.row, tt.col)
		if result != tt.expectedRune {
			t.Errorf("DisplayCell(%d, %d): got %c, want %c", tt.row, tt.col, result, tt.expectedRune)
		}
	}
}
