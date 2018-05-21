package players

import (
	"devshorts/logtest/game/boards"
	"devshorts/logtest/game/data"
)

type Player interface {
	GetMove(b *boards.Board) data.Position
	Token() data.Token
	Id() string
}
