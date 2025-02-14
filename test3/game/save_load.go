package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SaveGame(player1, player2 *Player, turn int, currentPlayer string) {
	fmt.Println("Сохранение игры...")

	// Сохраняем доски игроков
	saveBoardToFile("config/lastgame_player_1.cfg", player1.Board)
	saveBoardToFile("config/lastgame_player_2.cfg", player2.Board)

	// Сохраняем текущий ход и игрока
	file, err := os.Create("config/lastgame.cfg")
	if err != nil {
		fmt.Println("Ошибка сохранения состояния игры:", err)
		return
	}
	defer file.Close()

	// Пишем строго по формату
	_, err = file.WriteString(fmt.Sprintf("%d\n%s\n", turn, currentPlayer))
	if err != nil {
		fmt.Println("Ошибка записи в файл lastgame.cfg:", err)
	}
	fmt.Println("Игра сохранена!")
}

func saveBoardToFile(filename string, board Board) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Ошибка при сохранении в файл %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	// Сохраняем сетку поля
	for i, row := range board.Grid {
		for j, cell := range row {
			if board.Hits[i][j] {
				_, _ = file.WriteString("X")
			} else {
				_, _ = file.WriteString(string(cell))
			}
		}
		_, _ = file.WriteString("\n")
	}

	// Сохраняем корабли
	for _, ship := range board.Ships {
		line := ""
		for _, cell := range ship.Cells {
			line += fmt.Sprintf("%d,%d ", cell[0], cell[1])
		}
		_, _ = file.WriteString(strings.TrimSpace(line) + "\n")
	}
}

// LoadGame загружает сохранённую игру
func LoadGame() (*Player, *Player, int, string, error) {
	// Чтение файла с отладочным выводом
	fileContent, err := os.ReadFile("config/lastgame.cfg")
	if err != nil {
		return nil, nil, 0, "", fmt.Errorf("ошибка чтения файла lastgame.cfg: %v", err)
	}
	fmt.Println("Содержимое lastgame.cfg:")
	fmt.Println(string(fileContent))

	// Чтение и обработка
	if _, err := os.Stat("config/lastgame_player_1.cfg"); os.IsNotExist(err) {
		return nil, nil, 0, "", fmt.Errorf("файл lastgame_player_1.cfg не найден")
	}
	if _, err := os.Stat("config/lastgame_player_2.cfg"); os.IsNotExist(err) {
		return nil, nil, 0, "", fmt.Errorf("файл lastgame_player_2.cfg не найден")
	}
	if _, err := os.Stat("config/lastgame.cfg"); os.IsNotExist(err) {
		return nil, nil, 0, "", fmt.Errorf("файл lastgame.cfg не найден")
	}

	player1Board := LoadBoard("config/lastgame_player_1.cfg")
	player2Board := LoadBoard("config/lastgame_player_2.cfg")

	file, err := os.Open("config/lastgame.cfg")
	if err != nil {
		return nil, nil, 0, "", fmt.Errorf("ошибка открытия файла lastgame.cfg: %v", err)
	}
	defer file.Close()

	var turn int
	var currentPlayer string

	// Чтение с дополнительной обработкой ошибок
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		turn, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			return nil, nil, 0, "", fmt.Errorf("ошибка чтения номера хода: %v", err)
		}
	}

	if scanner.Scan() {
		currentPlayer = strings.TrimSpace(scanner.Text())
	} else {
		return nil, nil, 0, "", fmt.Errorf("ошибка чтения текущего игрока")
	}

	player1 := &Player{Name: "Player 1", Board: player1Board}
	player2 := &Player{Name: "Player 2", Board: player2Board}

	return player1, player2, turn, currentPlayer, nil
}

func LoadBoard(filename string) Board {
	var board Board
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Ошибка при загрузке файла %s: %v\n", filename, err)
		return board
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		if row < 10 {
			// Чтение игровой сетки
			for col, char := range line {
				switch char {
				case '1': // Корабль
					board.Grid[row][col] = '1'
				case '0': // Пустая клетка
					board.Grid[row][col] = '0'
				case 'X': // Попадание
					board.Grid[row][col] = '1'
					board.Hits[row][col] = true
				}
			}
			row++
		} else {
			// Чтение координат кораблей
			ship := Ship{}
			coords := strings.Split(line, " ")
			for _, coord := range coords {
				parts := strings.Split(coord, ",")
				if len(parts) != 2 {
					fmt.Printf("Ошибка: некорректные координаты '%s' в файле %s\n", coord, filename)
					continue
				}
				r, err1 := strconv.Atoi(parts[0])
				c, err2 := strconv.Atoi(parts[1])
				if err1 != nil || err2 != nil || r < 0 || r >= 10 || c < 0 || c >= 10 {
					fmt.Printf("Ошибка: некорректные координаты '%s' в файле %s\n", coord, filename)
					continue
				}
				ship.Cells = append(ship.Cells, [2]int{r, c})
			}
			board.Ships = append(board.Ships, &ship)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка при чтении файла %s: %v\n", filename, err)
	}

	board.ShipCount = len(board.Ships) // Обновляем количество кораблей
	return board
}

// countShips подсчитывает количество кораблей на доске
func countShips(board *Board) []*Ship {
	visited := [10][10]bool{}
	var ships []*Ship

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
				ship := &Ship{Cells: shipCells}
				ships = append(ships, ship)
			}
		}
	}

	return ships
}
