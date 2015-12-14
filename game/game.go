package game

import (
	"time"
)

var Pmanager PlayerManager
var Bmanager BulletManager

func StartGame() {
	go func() {
		timer := time.Tick(5 * time.Millisecond)
		for range timer {
			//Update game logic here..
			Pmanager.Update()
			Bmanager.Update()
		}
	}()
}
