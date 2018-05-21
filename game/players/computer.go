package players

import (
	"devshorts/logtest/game/data"
	"devshorts/logtest/game/boards"
	"math/rand"
)

type computerPlayer struct {
	token    data.Token
	name     string
	strategy func(b *boards.Board) data.Position
}

var _ Player = computerPlayer{}

func NewRandomComputer(token data.Token) Player {
	return computerPlayer{
		token:    token,
		name:     "Blorpus",
		strategy: randomStrategy,
	}
}

func randomStrategy(b *boards.Board) data.Position {
	ranPos := data.Position{
		X: rand.Intn(b.Size),
		Y: rand.Intn(b.Size),
	}

	if b.CanPlace(ranPos) {
		return ranPos
	}

	return randomStrategy(b)
}

func (computer computerPlayer) Token() data.Token {
	return computer.token
}

func (computer computerPlayer) Id() string {
	return computer.name
}

func (computer computerPlayer) GetMove(b *boards.Board) data.Position {
	return computer.strategy(b)
}
