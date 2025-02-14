package unit

import (
	"testing"

	"battleship/game"
)

func TestParseCoordinates(t *testing.T) {
	tests := []struct {
		input       string
		expectedRow int
		expectedCol int
		expectError bool
	}{
		{"a1", 0, 0, false},
		{"j10", 9, 9, false},
		{"A1", 0, 0, false}, // Верхний регистр
		{"b2", 1, 1, false},
		{"i10", 9, 8, false},
		{"j9", 8, 9, false},
		{"k1", 0, 0, true},  // Ввод за пределами поля
		{"a11", 0, 0, true}, // Некорректный ввод
		{"1a", 0, 0, true},  // Неверный формат
		{"", 0, 0, true},    // Пустой ввод
	}

	for _, tt := range tests {
		row, col, err := game.ParseCoordinates(tt.input)
		if (err != nil) != tt.expectError {
			t.Errorf("ParseCoordinates(%q) ошибка: ожидается %v, получено %v", tt.input, tt.expectError, err != nil)
		}
		if row != tt.expectedRow || col != tt.expectedCol {
			t.Errorf("ParseCoordinates(%q): ожидалось (%d, %d), получено (%d, %d)", tt.input, tt.expectedRow, tt.expectedCol, row, col)
		}
	}
}
