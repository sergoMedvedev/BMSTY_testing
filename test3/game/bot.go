package game

import (
	"fmt"
	"math/rand"
)

// Bot представляет бота-игрока
type Bot struct {
	Name    string
	Targets [][2]int // Очередь клеток для стрельбы по поврежденным кораблям
}

// MakeMove выполняет выстрел бота
func (b *Bot) MakeMove(opponent *Player) (int, int, bool) {
	var x, y int
	if len(b.Targets) > 0 {
		// Стреляем в первую клетку из очереди целей
		x, y = b.Targets[0][0], b.Targets[0][1]
		b.Targets = b.Targets[1:]
		fmt.Printf("%s стреляет в соседнюю клетку: (%d, %d)\n", b.Name, x+1, y+1)
	} else {
		// Выбираем случайную клетку, в которую еще не стреляли
		for {
			x, y = rand.Intn(10), rand.Intn(10)
			if !opponent.Board.Hits[x][y] {
				fmt.Printf("%s выбирает случайную клетку: (%d, %d)\n", b.Name, x+1, y+1)
				break
			}
		}
	}

	// Проверяем попадание
	hit := opponent.Board.Grid[x][y] == '1'
	if hit {
		fmt.Printf("%s попал в корабль!\n", b.Name)
		opponent.Board.Grid[x][y] = 'X'
		opponent.Board.ShipCount--

		// Добавляем соседние клетки в очередь целей
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				nx, ny := x+dx, y+dy
				if isValidCell(nx, ny) && opponent.Board.Grid[nx][ny] == '1' && !opponent.Board.Hits[nx][ny] {
					b.Targets = append(b.Targets, [2]int{nx, ny})
				}
			}
		}
	} else {
		fmt.Printf("%s промахнулся.\n", b.Name)
	}
	opponent.Board.Hits[x][y] = true

	return x, y, hit
}

// isValidCell проверяет, находится ли клетка в пределах игрового поля
func isValidCell(x, y int) bool {
	return x >= 0 && x < 10 && y >= 0 && y < 10
}
