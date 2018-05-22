package boards

import (
	"errors"
	"io"
	"fmt"
	"devshorts/logtest/game/data"
)

var noWinExists = errors.New("not a win")

type Board struct {
	Size   int
	tokens [][]data.Token
}

func NewBoard(size int) *Board {
	b := make([][]data.Token, size)
	for i := range b {
		b[i] = make([]data.Token, size)
	}

	return &Board{
		size,
		b,
	}
}

func (b *Board) CanPlace(p data.Position) bool {
	return !b.get(p.X, p.Y).IsSet()
}

func (b *Board) Print(w io.Writer) {
	for i := range b.tokens {
		for j := range b.tokens {
			switch b.tokens[i][j] {
			case data.X:
				fmt.Fprint(w, "X")
			case data.O:
				fmt.Fprint(w, "O")
			case data.Unset:
				fmt.Fprint(w, "-")
			}
		}

		fmt.Fprintln(w)
	}
}

func (b *Board) PutToken(token data.Token, position data.Position) error {
	if b.CanPlace(position) {
		b.tokens[position.X][position.Y] = token
	} else {
		return errors.New("cannot place position since it is taken")
	}

	return nil
}

func (b *Board) colWinner(col int) (*data.Token, error) {
	prevPlayer := b.get(col, 0)

	if !prevPlayer.IsSet() {
		return nil, noWinExists
	}

	for _, p := range b.tokens[col] {
		if p != prevPlayer || !p.IsSet() {
			return nil, noWinExists
		}
	}

	return &prevPlayer, nil
}

func (b *Board) rowWinner(row int) (*data.Token, error) {
	prevPlayer := b.get(0, row)

	if !prevPlayer.IsSet() {
		return nil, noWinExists
	}

	for i := range b.tokens {
		currentPlayer := b.get(i, row)
		if currentPlayer != prevPlayer || !currentPlayer.IsSet() {
			return nil, noWinExists
		}
	}

	return &prevPlayer, nil
}

func (b *Board) diagonalWinner() (*data.Token, error) {
	rightToLeft := func() (*data.Token, error) {
		prevPlayer := b.get(0, b.Size-1)

		if !prevPlayer.IsSet() {
			return nil, noWinExists
		}

		for i := range b.tokens {
			if b.get(i, b.Size-i-1) != prevPlayer {
				return nil, noWinExists
			}
		}

		return &prevPlayer, nil
	}

	prevPlayer := b.tokens[0][0]

	for i := range b.tokens {
		if b.get(i, i) != prevPlayer {
			winner, err := rightToLeft()
			if err == nil {
				return winner, nil
			} else {
				return nil, noWinExists
			}
		}
	}

	if !prevPlayer.IsSet() {
		return nil, noWinExists
	}

	return &prevPlayer, nil
}

func (b *Board) get(x, y int) data.Token {
	return b.tokens[x][y]
}

func (b *Board) IsGameOver() bool {
	_, winnerExists := b.GetWinner()
	if winnerExists {
		return true
	}

	for i := range b.tokens {
		for j := range b.tokens {
			if !b.tokens[i][j].IsSet() {
				return false
			}
		}
	}

	return true
}

func (b *Board) GetWinner() (*data.Token, bool) {
	diagonalWinner, err := b.diagonalWinner()
	if err == nil {
		return diagonalWinner, true
	}

	for i := range b.tokens {
		colWinner, err := b.colWinner(i)
		if err == nil {
			return colWinner, true
		}

		rowWinner, err := b.rowWinner(i)
		if err == nil {
			return rowWinner, true
		}
	}

	return nil, false
}
