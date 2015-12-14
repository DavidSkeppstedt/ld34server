package player

import (
	"math"
)

type Bullet struct {
	Pos   Position `json:"position"`
	Angle float32  `json:"angle"`
	Speed float32  `json:speed"`
	Dead  bool
}

func (this *Bullet) move() {

	this.Pos.X += speed * float32(math.Cos(float64(this.Angle)*math.Pi/180))
	this.Pos.Y += speed * float32(math.Sin(float64(this.Angle)*math.Pi/180))
}
func (this *Bullet) outside() {
	if this.Pos.X > float32(width) || this.Pos.X < 0 {
		this.Dead = true
		return
	}
	if this.Pos.Y > float32(heigth) || this.Pos.Y < 0 {
		this.Dead = true
		return
	}
	this.Dead = false
}
func (this *Bullet) Update() {
	this.move()
	this.outside()
}

type BulletManager struct {
	Bullets []Bullet
}

func (this *BulletManager) NewBullet(player Player) {
	bullet := &Bullet{player.Pos, player.Angle, 5, true}
	this.Bullets = append(this.Bullets, *bullet)
}
func (this *BulletManager) Update() {
	for i, bullet := range this.Bullets {
		bullet.Update()
		if bullet.Dead {
			this.Bullets = append(this.Bullets[:i], this.Bullets[i+1:]...)
		}
	}
}
