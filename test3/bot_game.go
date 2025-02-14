package main

import (
	"battleship/game"
	"fmt"
	"os"
)

// countShips подсчитывает количество кораблей на доске
func countShips(board *game.Board) []*game.Ship {
	visited := [10][10]bool{}
	var ships []*game.Ship

	// Смещения для поиска соседей
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// Рекурсивная функция DFS для поиска ячеек корабля
	var dfs func(row, col int, shipCells *[][2]int)
	dfs = func(row, col int, shipCells *[][2]int) {
		visited[row][col] = true
		*shipCells = append(*shipCells, [2]int{row, col})

		for _, dir := range directions {
			nr, nc := row+dir[0], col+dir[1]
			if nr >= 0 && nr < 10 && nc >= 0 && nc < 10 && !visited[nr][nc] && board.Grid[nr][nc] == '1' {
				dfs(nr, nc, shipCells)
			}
		}
	}

	// Основной цикл по клеткам
	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			if board.Grid[row][col] == '1' && !visited[row][col] {
				var shipCells [][2]int
				dfs(row, col, &shipCells)

				// Создаём новый корабль
				ship := &game.Ship{Cells: shipCells}
				ships = append(ships, ship)
			}
		}
	}

	return ships
}

func main() {
	fmt.Println("=== Состояние начальных досок ===")
	player1 := &game.Player{Name: "Bot 1", Board: game.LoadBoard("config/newgame_player_1.cfg")}
	player2 := &game.Player{Name: "Bot 2", Board: game.LoadBoard("config/newgame_player_2.cfg")}
	bot1 := &game.Bot{Name: "Bot 1"}
	bot2 := &game.Bot{Name: "Bot 2"}

	// Подсчёт кораблей после загрузки
	player1.Board.Ships = countShips(&player1.Board)
	player2.Board.Ships = countShips(&player2.Board)
	player1.Board.ShipCount = len(player1.Board.Ships)
	player2.Board.ShipCount = len(player2.Board.Ships)

	fmt.Printf("%s: %d кораблей\n", player1.Name, player1.Board.ShipCount)
	fmt.Printf("%s: %d кораблей\n", player2.Name, player2.Board.ShipCount)

	fmt.Printf("%s (ваше поле):\n", player1.Name)
	game.DisplayBoards(player1.Board, player2.Board)
	fmt.Printf("%s (ваше поле):\n", player2.Name)
	game.DisplayBoards(player2.Board, player1.Board)

	fmt.Println("=== Начало игры ===")

	activePlayer := player1
	activeBot := bot1
	opponent := player2
	opponentBot := bot2
	turn := 1
	halfShipWarning := map[*game.Player]bool{player1: false, player2: false}
	initialShipCount := map[*game.Player]int{
		player1: player1.Board.ShipCount,
		player2: player2.Board.ShipCount,
	}

	// Основной цикл игры
	for player1.Board.ShipCount > 0 && player2.Board.ShipCount > 0 {
		fmt.Printf("\nХод %d. Ходит: %s\n", turn, activeBot.Name)

		// Выполняем выстрел
		x, y, hit := activeBot.MakeMove(opponent)
		coords := formatCoordinates(x, y)
		fmt.Printf("%s стреляет в %s: %s\n", activeBot.Name, coords, map[bool]string{true: "Попадание", false: "Мимо"}[hit])

		// Обновление состояния кораблей
		if hit {
			for _, ship := range opponent.Board.Ships {
				if ship.MarkHit(x, y) && ship.IsSunk {
					fmt.Printf("Корабль %s затоплен!\n", opponent.Name)
					opponent.Board.ShipCount--
				}
			}
		}

		// Проверка на "половину затопленных кораблей"
		if opponent.Board.ShipCount*2 <= initialShipCount[opponent] && !halfShipWarning[opponent] {
			halfShipWarning[opponent] = true
			fmt.Printf("\nПоловина кораблей %s затоплена! Игра приостановлена.\n", opponent.Name)
			saveGame(player1, player2, turn, activePlayer.Name)
			if !handlePause() {
				fmt.Println("Игра завершена командой 'exit'.")
				os.Exit(0)
			}
		}

		// Если выстрел мимо, меняем игрока
		if !hit {
			activePlayer, opponent = opponent, activePlayer
			activeBot, opponentBot = opponentBot, activeBot
			turn++
		}

		// Вывод состояния досок
		fmt.Println("\nТекущее состояние игры:")
		fmt.Printf("%s (ваше поле):\n", activePlayer.Name)
		game.DisplayBoards(activePlayer.Board, opponent.Board)
	}

	// Завершение игры
	if player1.Board.ShipCount == 0 {
		fmt.Println("Bot 2 победил!")
	} else {
		fmt.Println("Bot 1 победил!")
	}
	fmt.Println("=== Игра завершена ===")
}

// saveGame сохраняет состояние игры в файлы
func saveGame(player1, player2 *game.Player, turn int, currentPlayer string) {
	fmt.Println("Сохраняем игру...")
	game.SaveGame(player1, player2, turn, currentPlayer)
	fmt.Println("Игра сохранена!")
}

// handlePause обрабатывает паузу игры, ожидая ввода "continue" или "exit"
func handlePause() bool {
	fmt.Println("Введите 'continue', чтобы продолжить игру, или 'exit', чтобы выйти.")
	var input string
	for {
		fmt.Print("> ")
		fmt.Scanln(&input)
		if input == "continue" {
			fmt.Println("Игра продолжается!")
			return true
		}
		if input == "exit" {
			fmt.Println("Завершаем игру...")
			return false
		}
		fmt.Println("Неверная команда. Введите 'continue' или 'exit'.")
	}
}

// formatCoordinates преобразует индексы строки и столбца в формат "(буква, число)"
func formatCoordinates(row, col int) string {
	columnLetter := string(rune('A' + col))             // Преобразование столбца в букву
	return fmt.Sprintf("(%s, %d)", columnLetter, row+1) // Строка начинается с 1
}
