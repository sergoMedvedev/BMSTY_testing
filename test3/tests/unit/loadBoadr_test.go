package unit

import (
	"battleship/game"
	"testing"
)

func TestLoadBoard(t *testing.T) {
	expectedGrid := [10][10]rune{
		{'1', '0', '0'},
		{'0', '1', '0'},
		{'0', '0', '1'},
	}

	board := game.LoadBoard("test_board.cfg")

	for i, row := range expectedGrid {
		for j, cell := range row {
			if board.Grid[i][j] != cell {
				t.Errorf("LoadBoard: cell (%d, %d): got %c, want %c", i, j, board.Grid[i][j], cell)
			}
		}
	}
}
