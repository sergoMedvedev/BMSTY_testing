package game

// DisplayCell возвращает отображение своей клетки (своё поле)
func (b *Board) DisplayCell(row, col int) rune {
	if b.Hits[row][col] { // Клетка была атакована
		return 'X' // Все выстрелы обозначаются как X
	}
	return b.Grid[row][col] // '1' (корабль) или '0' (пустая клетка)
}

// HiddenCell возвращает отображение вражеской клетки (поле противника)
func (b *Board) HiddenCell(row, col int) rune {
	if b.Hits[row][col] { // Клетка была атакована
		return 'X' // Все выстрелы обозначаются как X
	}
	return b.Grid[row][col] // '1' (корабль) или '0' (пустая клетка)
}

// MarkHit отмечает попадание в корабль и проверяет его состояние
func (s *Ship) MarkHit(x, y int) bool {
	for _, cell := range s.Cells {
		if cell[0] == x && cell[1] == y {
			s.HitCount++
			if s.HitCount == len(s.Cells) {
				s.IsSunk = true
			}
			return true
		}
	}
	return false
}
