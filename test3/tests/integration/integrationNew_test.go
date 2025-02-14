package integration

import (
	"battleship/game"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// BotPlayer представляет бота-игрока
type BotPlayer struct {
	Name           string
	Board          game.Board
	AvailableShots map[string]bool // Доступные для стрельбы клетки
}

// NewBotPlayer создает нового бота
func NewBotPlayer(name string, board game.Board) *BotPlayer {
	availableShots := make(map[string]bool)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			availableShots[fmt.Sprintf("%d,%d", i, j)] = true
		}
	}
	return &BotPlayer{Name: name, Board: board, AvailableShots: availableShots}
}

// ChooseShot выбирает клетку для выстрела
func (b *BotPlayer) ChooseShot() (int, int) {
	shots := make([]string, 0, len(b.AvailableShots))
	for shot := range b.AvailableShots {
		shots = append(shots, shot)
	}
	chosen := shots[rand.Intn(len(shots))]
	var row, col int
	fmt.Sscanf(chosen, "%d,%d", &row, &col)
	return row, col
}

// RegisterShot удаляет клетку из доступных для стрельбы
func (b *BotPlayer) RegisterShot(row, col int) {
	delete(b.AvailableShots, fmt.Sprintf("%d,%d", row, col))
}

// DisplayBoards отображает текущие доски
func DisplayBoards(player1, player2 *BotPlayer) {
	fmt.Println("\n--- Текущее состояние игры ---")
	fmt.Printf("%s (ваше поле):\n", player1.Name)
	game.DisplayBoards(player1.Board, player2.Board)
	fmt.Printf("\n%s (ваше поле):\n", player2.Name)
	game.DisplayBoards(player2.Board, player1.Board)
	fmt.Println("--------------------------------")
}

// TestGameIntegration реализует интеграционный тест
func TestGameIntegration(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// Создаем двух ботов
	player1 := NewBotPlayer("Player 1", game.LoadBoard("config/newgame_player_1.cfg"))
	player2 := NewBotPlayer("Player 2", game.LoadBoard("config/newgame_player_2.cfg"))

	activePlayer, opponent := player1, player2
	turns := 0

	for {
		turns++
		fmt.Printf("\n%s: ход №%d\n", activePlayer.Name, turns)

		// Показываем доски перед ходом
		DisplayBoards(player1, player2)

		// Выбор выстрела
		row, col := activePlayer.ChooseShot()
		activePlayer.RegisterShot(row, col)

		// Выполнение выстрела
		if opponent.Board.Hits[row][col] {
			fmt.Printf("Попадание в клетку (%d, %d)!\n", row+1, col+1)
			opponent.Board.Hits[row][col] = true
			if opponent.Board.ShipCount == 0 {
				fmt.Printf("\n%s победил! Игра завершена.\n", activePlayer.Name)
				break
			}
			continue
		} else {
			fmt.Printf("Мимо! Выстрел в клетку (%d, %d).\n", row+1, col+1)
		}

		// Проверка завершения игры при уничтожении половины кораблей
		if opponent.Board.ShipCount <= opponent.Board.ShipCount/2 {
			fmt.Printf("\nИгра остановлена: половина кораблей %s уничтожена.\n", opponent.Name)
			break
		}

		// Передача хода
		activePlayer, opponent = opponent, activePlayer
	}

	fmt.Println("\nИгра завершена!")
}
