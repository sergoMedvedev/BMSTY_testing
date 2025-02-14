package game

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBotGame(t *testing.T) {
	// Инициализация игровых досок для двух ботов
	player1 := &Player{Name: "Bot_t 1", Board: LoadBoard("../config/newgame_player_1.cfg")}
	player2 := &Player{Name: "Bot_t 2", Board: LoadBoard("../config/newgame_player_2.cfg")}

	// Подсчёт кораблей после загрузки досок
	player1.Board.Ships = countShips(&player1.Board)
	player2.Board.Ships = countShips(&player2.Board)
	player1.Board.ShipCount = len(player1.Board.Ships)
	player2.Board.ShipCount = len(player2.Board.Ships)

	bot1 := createBot(player1)
	bot2 := createBot(player2)

	fmt.Printf("Начало игры: %s с %d кораблями, %s с %d кораблями\n",
		player1.Name, player1.Board.ShipCount, player2.Name, player2.Board.ShipCount)

	turn := 1
	activeBot := bot1
	opponentBot := bot2
	for player1.Board.ShipCount > 0 && player2.Board.ShipCount > 0 {
		fmt.Printf("Ход %d. Ходит: %s\n", turn, activeBot.Player.Name)

		x, y, hit := activeBot.makeMove(opponentBot.Player)
		coords := formatCoordinates(x, y)
		fmt.Printf("%s стреляет в %s: %s\n", activeBot.Player.Name, coords, map[bool]string{true: "Попадание", false: "Мимо"}[hit])

		// Проверка потопления корабля
		if hit {
			for _, ship := range opponentBot.Player.Board.Ships {
				if ship.MarkHit(x, y) && ship.IsSunk {
					fmt.Printf("Корабль %s уничтожен! Осталось кораблей: %d\n", opponentBot.Player.Name, opponentBot.Player.Board.ShipCount-1)
					opponentBot.Player.Board.ShipCount--
				}
			}
		}

		// Проверка завершения игры
		if opponentBot.Player.Board.ShipCount == 0 {
			fmt.Printf("%s победил!\n", activeBot.Player.Name)
			return
		}

		// Смена хода
		if !hit {
			activeBot, opponentBot = opponentBot, activeBot
			turn++
		}
	}
}

// Bot_t представляет игрока-бота
type Bot_t struct {
	Player *Player
	Shots  map[[2]int]bool // Клетки, в которые уже стреляли
}

// createBot создает бота
func createBot(player *Player) *Bot_t {
	return &Bot_t{
		Player: player,
		Shots:  make(map[[2]int]bool),
	}
}

// makeMove выполняет ход бота
func (b *Bot_t) makeMove(opponent *Player) (int, int, bool) {
	// Поиск поврежденного, но не потопленного корабля
	for _, ship := range opponent.Board.Ships {
		if !ship.IsSunk {
			for _, cell := range ship.Cells {
				if opponent.Board.Hits[cell[0]][cell[1]] {
					// Стреляем по соседним клеткам
					directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
					for _, dir := range directions {
						x, y := cell[0]+dir[0], cell[1]+dir[1]
						if isValidMove(x, y, b) {
							b.Shots[[2]int{x, y}] = true
							hit := opponent.Board.Grid[x][y] == '1'
							opponent.Board.Hits[x][y] = true
							return x, y, hit
						}
					}
				}
			}
		}
	}

	// Если нет поврежденных кораблей, стреляем в случайную клетку
	for {
		x, y := rand.Intn(10), rand.Intn(10)
		if isValidMove(x, y, b) {
			b.Shots[[2]int{x, y}] = true
			hit := opponent.Board.Grid[x][y] == '1'
			opponent.Board.Hits[x][y] = true
			return x, y, hit
		}
	}
}

// isValidMove проверяет, можно ли стрелять в клетку
func isValidMove(x, y int, b *Bot_t) bool {
	return x >= 0 && x < 10 && y >= 0 && y < 10 && !b.Shots[[2]int{x, y}]
}

// formatCoordinates преобразует индексы строки и столбца в формат "(буква, число)"
func formatCoordinates(row, col int) string {
	columnLetter := string(rune('A' + col))
	return fmt.Sprintf("(%s, %d)", columnLetter, row+1)
}
