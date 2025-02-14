package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// StartNewGame начинает новую игру
func StartNewGame() {
	player1 := Player{Name: "Player 1", Board: LoadBoard("config/newgame_player_1.cfg")}
	player2 := Player{Name: "Player 2", Board: LoadBoard("config/newgame_player_2.cfg")}

	// Пересчитываем корабли после загрузки досок
	player1.Board.Ships = countShips(&player1.Board)
	player1.Board.ShipCount = len(player1.Board.Ships)

	player2.Board.Ships = countShips(&player2.Board)
	player2.Board.ShipCount = len(player2.Board.Ships)

	fmt.Println("Игра началась! Первый ход за Player 1.")
	RunGame(&player1, &player2, 1, "Player 1")
}

// RunGame управляет ходами игроков
func RunGame(player1, player2 *Player, turn int, currentPlayer string) {
	var activePlayer, opponent *Player
	if currentPlayer == "Player 1" {
		activePlayer = player1
		opponent = player2
	} else {
		activePlayer = player2
		opponent = player1
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s, ваш ход:\n", activePlayer.Name)
		DisplayBoards(activePlayer.Board, opponent.Board)
		fmt.Print("Введите координаты (например, a1) или 'exit' для сохранения: ")

		// Чтение ввода
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			// Сохраняем игру перед выходом
			SaveGame(player1, player2, turn, activePlayer.Name)
			fmt.Println("Игра сохранена. До свидания!")
			return // Выход из функции
		}

		// Обработка ввода
		x, y, err := ParseCoordinates(input)
		if err != nil {
			fmt.Println("Ошибка: Неверный формат ввода. Попробуйте снова.")
			continue
		}

		// Проверка повторного выстрела
		if opponent.Board.Hits[x][y] {
			fmt.Println("Вы уже стреляли в эту клетку. Попробуйте снова.")
			continue
		}

		// Обработка выстрела
		if activePlayer.MakeMove(x, y, opponent) {
			if opponent.Board.ShipCount == 0 {
				fmt.Printf("Поздравляем, %s победил!\n", activePlayer.Name)
				return // Игра завершается
			}
			fmt.Println("Попадание! Вы можете стрелять еще раз.")
		} else {
			fmt.Println("Мимо! Ход переходит другому игроку.")
			// Переход хода
			activePlayer, opponent = opponent, activePlayer
			turn++
		}
	}
}

// DisplayBoards отображает статус кораблей: свои слева, чужие справа
func DisplayBoards(ownBoard, enemyBoard Board) {
	fmt.Println("|    Ваше поле:         |      |  Поле противника:    |")
	fmt.Println("   A B C D E F G H I J            A B C D E F G H I J")

	for i := 0; i < 10; i++ {
		fmt.Printf("%2d ", i+1)
		for j := 0; j < 10; j++ {
			fmt.Print(string(ownBoard.DisplayCell(i, j)), " ")
		}
		fmt.Printf("        %2d ", i+1)
		for j := 0; j < 10; j++ {
			fmt.Print(string(enemyBoard.HiddenCell(i, j)), " ")
		}
		fmt.Println()
	}
}

// ParseCoordinates преобразует строку координат в индексы x и y
func ParseCoordinates(input string) (int, int, error) {
	// Проверяем длину ввода: минимум 2 символа (например, a1), максимум 3 (например, j10)
	if len(input) < 2 || len(input) > 3 {
		return 0, 0, fmt.Errorf("некорректный ввод")
	}

	// Первая часть ввода — буква (столбец)
	letter := strings.ToLower(string(input[0]))
	col := int(letter[0] - 'a') // Преобразование буквы в индекс столбца

	// Вторая часть ввода — цифра (строка)
	number := input[1:]
	row, err := strconv.Atoi(number)
	if err != nil || row < 1 || row > 10 || col < 0 || col >= 10 {
		return 0, 0, fmt.Errorf("некорректный ввод")
	}

	// Преобразование в индексы массива (0-based)
	rowIndex := row - 1

	// Отладочный вывод
	fmt.Printf("Ввод: %s -> Строка: %d, Столбец: %d\n", input, rowIndex, col)

	return rowIndex, col, nil
}

// ContinueGame продолжает сохранённую игру
func ContinueGame() {
	player1, player2, turn, currentPlayer, err := LoadGame()
	if err != nil {
		fmt.Println("Ошибка загрузки сохраненной игры:", err)
		return
	}

	fmt.Println("Игра успешно загружена. Продолжаем!")
	RunGame(player1, player2, turn, currentPlayer)
}
