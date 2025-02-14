package game

type Board struct {
	Grid      [10][10]rune // '1' — часть корабля, '0' — пустая клетка, 'X' — попадание
	Hits      [10][10]bool // Отмеченные выстрелы
	Ships     []*Ship      // Список кораблей
	ShipCount int          // Количество оставшихся кораблей
}

type Ship struct {
	Cells    [][2]int // Координаты ячеек корабля
	IsSunk   bool     // Затоплен ли корабль
	HitCount int      // Количество попаданий
}

type Player struct {
	Name  string
	Board Board
}
