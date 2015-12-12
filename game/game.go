package game

import (
	"github.com/DavidSkeppstedt/ld34server/game/player"
	_ "time"
)

var Pmanager player.PlayerManager

func StartGame() {
	go func() {
		/*timer := time.Tick(5 * time.Millisecond)
		for now := range timer {
			//Update game logic here..
			log.Println(now)
		}*/
	}()
}
