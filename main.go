package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"github.com/onrik/logrus/filename"
	"devshorts/logtest/game"
	"devshorts/logtest/game/players"
	"devshorts/logtest/game/data"
)

func init() {
	filenameHook := filename.NewHook()

	log.AddHook(filenameHook)

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	gamePlayers := []players.Player{
		players.NewHumanPlayer(data.X, "Anton"),
		players.NewRandomComputer(data.O),
	}

	game.Play(gamePlayers, 3)
}
