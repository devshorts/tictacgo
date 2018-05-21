package players

import (
	"devshorts/logtest/game/data"
	"devshorts/logtest/game/boards"
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

type humanPlayer struct {
	token data.Token
	name  string
}

var _ Player = humanPlayer{}

func NewHumanPlayer(token data.Token, name string) Player {
	return humanPlayer{token, name}
}

func (human humanPlayer) Token() data.Token {
	return human.token
}

func (human humanPlayer) Id() string {
	return human.name
}

func (human humanPlayer) GetMove(b *boards.Board) data.Position {
	fmt.Printf("Player %s make your move\n", human.Id())

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("X: ")

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	x, _ := strconv.Atoi(line)

	fmt.Print("Y: ")

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	y, _ := strconv.Atoi(line)

	pos := data.Position{x, y}
	if !b.CanPlace(pos) {
		fmt.Println("Cannot place at position, try again")

		return human.GetMove(b)
	}

	return pos
}

