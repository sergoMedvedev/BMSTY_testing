package unit

import (
	"battleship/game"
	"testing"
)

func TestMakeMove(t *testing.T) {
	player := game.Player{
		Board: game.Board{
			Grid: [10][10]rune{
				{'1', '0', '1'},
			},
			Hits: [10][10]bool{},
		},
	}

	tests := []struct {
		row, col       int
		expectedHit    bool
		expectedHitVal rune
	}{
		{0, 0, true, 'X'},  // Попадание
		{0, 1, false, 'X'}, // Промах
		{0, 2, true, 'X'},  // Попадание
	}

	for _, tt := range tests {
		result := player.MakeMove(tt.row, tt.col, &player)
		if result != tt.expectedHit || player.Board.Hits[tt.row][tt.col] != true || player.Board.Grid[tt.row][tt.col] != tt.expectedHitVal {
			t.Errorf("MakeMove(%d, %d): got %v, want %v", tt.row, tt.col, result, tt.expectedHit)
		}
	}
}
