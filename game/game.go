package game

import (
	"github.com/DavidSkeppstedt/ld34server/game/player"
	"time"
)

var Pmanager player.PlayerManager

func StartGame() {
	go func() {
		timer := time.Tick(5 * time.Millisecond)
		for range timer {
			//Update game logic here..
			Pmanager.Update()
		}
	}()
}
