package main

import (
	"github.com/DavidSkeppstedt/ld34server/game"
	"github.com/DavidSkeppstedt/ld34server/host"
)

func main() {
	game.StartGame()
	host.ListenAndServe()
}
