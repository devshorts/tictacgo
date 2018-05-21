package boards

import (
	"testing"
	"os"
	"devshorts/logtest/game/data"
)

func TestBoard_IsWin(t *testing.T) {
	board := NewBoard(3)

	_, winnerExists := board.GetWinner()

	if winnerExists {
		t.Fail()
	}

	board.Print(os.Stdout)
}

func TestBoard_ColWins(t *testing.T) {
	board := NewBoard(3)

	for i := 0; i < board.Size; i++ {
		board.PutToken(data.X, data.Position{
			X: i,
			Y: 0,
		})
	}

	token, _ := board.GetWinner()

	if *token != data.X {
		t.Fail()
	}

	board.Print(os.Stdout)
}

func TestBoard_RowWins(t *testing.T) {
	board := NewBoard(3)

	for i := 0; i < board.Size; i++ {
		board.PutToken(data.X, data.Position{
			X: 0,
			Y: i,
		})
	}

	token, _ := board.GetWinner()

	if *token != data.X {
		t.Fail()
	}

	board.Print(os.Stdout)
}

func TestBoard_DiagonalWins(t *testing.T) {
	board := NewBoard(3)

	for i := 0; i < board.Size; i++ {
		board.PutToken(data.X, data.Position{
			X: i,
			Y: i,
		})
	}

	token, _ := board.GetWinner()

	if *token != data.X {
		t.Fail()
	}

	board.Print(os.Stdout)
}
