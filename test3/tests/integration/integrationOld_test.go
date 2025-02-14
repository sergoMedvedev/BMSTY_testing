package integration

import (
	"battleship/game"
	"fmt"
	"testing"
)

func TestIntegrationGame(t *testing.T) {
	// Инициализация ботов и игрового поля
	player1 := &game.Player{Name: "Bot 1", Board: game.LoadBoard("../../config/newgame_player_1.cfg")}
	player2 := &game.Player{Name: "Bot 2", Board: game.LoadBoard("../../config/newgame_player_2.cfg")}
	bot1 := &game.Bot{Name: "Bot 1"}
	bot2 := &game.Bot{Name: "Bot 2"}

	activePlayer := player1
	activeBot := bot1
	opponent := player2
	opponentBot := bot2

	turn := 1
	for player1.Board.ShipCount > 0 && player2.Board.ShipCount > 0 {
		fmt.Printf("Ход %d. Ходит: %s\n", turn, activeBot.Name)

		// Выполняем выстрел
		x, y, hit := activeBot.MakeMove(opponent)
		fmt.Printf("%s стреляет в (%d, %d): %s\n", activeBot.Name, x+1, y+1, map[bool]string{true: "Попадание", false: "Мимо"}[hit])

		// Вывод состояния досок
		fmt.Println("\nТекущее состояние:")
		fmt.Printf("%s (ваше поле):\n", activePlayer.Name)
		game.DisplayBoards(activePlayer.Board, opponent.Board)
		fmt.Printf("%s (ваше поле):\n", opponent.Name)
		game.DisplayBoards(opponent.Board, activePlayer.Board)

		if !hit {
			// Передаем ход
			activePlayer, opponent = opponent, activePlayer
			activeBot, opponentBot = opponentBot, activeBot
			turn++
		}

		// Проверяем завершение игры при уничтожении половины кораблей
		if opponent.Board.ShipCount <= 10 {
			fmt.Printf("\nОдин из ботов потерял половину кораблей. Сохраняем и перезапускаем игру...\n")
			game.SaveGame(player1, player2, turn, activePlayer.Name)
			break
		}
	}

	// Завершение игры или продолжение после сохранения
	if player1.Board.ShipCount == 0 {
		fmt.Println("Bot 2 победил!")
	} else if player2.Board.ShipCount == 0 {
		fmt.Println("Bot 1 победил!")
	} else {
		fmt.Println("\nПродолжаем сохраненную игру...")

		player1, player2, turn, currentPlayer, err := game.LoadGame()
		if err != nil {
			t.Fatalf("Ошибка загрузки игры: %v", err)
		}

		if currentPlayer == "Bot 1" {
			activePlayer, opponent = player1, player2
			activeBot, opponentBot = bot1, bot2
		} else {
			activePlayer, opponent = player2, player1
			activeBot, opponentBot = bot2, bot1
		}

		for player1.Board.ShipCount > 0 && player2.Board.ShipCount > 0 {
			fmt.Printf("Ход %d. Ходит: %s\n", turn, activeBot.Name)

			x, y, hit := activeBot.MakeMove(opponent)
			fmt.Printf("%s стреляет в (%d, %d): %s\n", activeBot.Name, x+1, y+1, map[bool]string{true: "Попадание", false: "Мимо"}[hit])

			// Вывод состояния досок
			fmt.Println("\nТекущее состояние:")
			fmt.Printf("%s (ваше поле):\n", activePlayer.Name)
			game.DisplayBoards(activePlayer.Board, opponent.Board)
			fmt.Printf("%s (ваше поле):\n", opponent.Name)
			game.DisplayBoards(opponent.Board, activePlayer.Board)

			if !hit {
				activePlayer, opponent = opponent, activePlayer
				activeBot, opponentBot = opponentBot, activeBot
				turn++
			}
		}

		if player1.Board.ShipCount == 0 {
			fmt.Println("Bot 2 победил!")
		} else {
			fmt.Println("Bot 1 победил!")
		}
	}
}
