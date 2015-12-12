package game

import (
	"time"
)

func StartGame() {
	go func() {
		timer := time.Tick(5 * Time.Millisecond)
		for now := range timer {
			//Update game logic here..
		}
	}()
}
