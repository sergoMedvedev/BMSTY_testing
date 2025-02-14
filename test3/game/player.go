package game

import "fmt"

// MakeMove обрабатывает выстрел игрока
func (p *Player) MakeMove(x, y int, opponent *Player) bool {
	if x < 0 || x >= 10 || y < 0 || y >= 10 {
		fmt.Println("Координаты за пределами поля. Попробуйте снова.")
		return false
	}

	if opponent.Board.Hits[x][y] {
		fmt.Println("Вы уже стреляли в эту клетку. Попробуйте снова.")
		return false
	}

	opponent.Board.Hits[x][y] = true
	if opponent.Board.Grid[x][y] == '1' {
		opponent.Board.Grid[x][y] = 'X'
		for _, ship := range opponent.Board.Ships {
			if ship.MarkHit(x, y) { // Обновляем состояние корабля
				if ship.IsSunk {
					fmt.Printf("Корабль противника уничтожен! Осталось кораблей: %d\n", opponent.Board.ShipCount-1)
					opponent.Board.ShipCount--
				}
				break
			}
		}
		return true
	}

	return false
}
