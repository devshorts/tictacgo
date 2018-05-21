package game

import (
	"devshorts/logtest/game/players"
	"devshorts/logtest/game/boards"
	"os"
	"fmt"
	"devshorts/logtest/game/data"
)

func Play(players []players.Player, boardSize int) {
	board := boards.NewBoard(boardSize)

	printBoard := func() {
		board.Print(os.Stdout)
	}

	for !board.IsGameOver() {
		for _, player := range players {
			printBoard()

			board.PutToken(player.Token(), player.GetMove(board))

			if board.IsGameOver() {
				break
			}
		}
	}

	winningToken, winnerExists := board.GetWinner()

	if winnerExists {
		fmt.Println("Congradulations to player " + findPlayer(players, winningToken))
	} else {
		fmt.Println("Statemate!")
	}
}

func findPlayer(players []players.Player, token *data.Token) string {
	for _, player := range players {
		if player.Token() == *token {
			return player.Id()
		}
	}

	return "unknown"
}
