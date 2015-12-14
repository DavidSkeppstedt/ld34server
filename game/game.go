package game

import (
	"time"
)

var Pmanager PlayerManager
var Bmanager BulletManager

func StartGame() {
	go func() {
		timer := time.Tick(10 * time.Millisecond)
		for range timer {
			//Update game logic here..
			Pmanager.Update()
			Bmanager.Update()
			for _, bullet := range Bmanager.Bullets {
				for _, player := range Pmanager.Players {
					if bullet.Pos.X > player.Pos.X-float32(player.Width)/2 && bullet.Pos.X < player.Pos.X+float32(player.Width)/2 {

						if bullet.Pos.Y > player.Pos.Y-float32(player.Heigth)/2 && bullet.Pos.Y < player.Pos.Y+float32(player.Heigth)/2 {
							bullet.Alive = false
							player.Alive = false
						}
					}
				}
			}
		}
	}()
}
