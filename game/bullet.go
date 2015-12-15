package game

import (
	"math"
)

type Bullet struct {
	Pos    Position `json:"position"`
	Angle  float32  `json:"angle"`
	Speed  float32  `json:"-"`
	Width  int      `json:"-"`
	Heigth int      `json:"-"`
	Dead   bool     `json:"-"`
	Alive  bool     `json:"-"`
}

func (this *Bullet) move() {

	this.Pos.X += this.Speed * float32(math.Cos(float64(this.Angle)*math.Pi/180))
	this.Pos.Y += this.Speed * float32(math.Sin(float64(this.Angle)*math.Pi/180))
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
	if this.Alive {
		this.move()
		this.outside()
	} else {
		this.Dead = true
	}
}

type BulletManager struct {
	Bullets []*Bullet
}

func (this *BulletManager) Get() []Bullet {
	var allBullets []Bullet
	if len(this.Bullets) < 1 {
		return []Bullet{}
	}
	for _, v := range this.Bullets {
		allBullets = append(allBullets, *v)
	}
	return allBullets
}

func (this *BulletManager) NewBullet(player Player) {
	bullet := &Bullet{player.Pos, player.Angle, 5, 2, 5, false, true}
	bullet.move()
	bullet.move()
	bullet.move()
	bullet.move()
	this.Bullets = append(this.Bullets, bullet)
}
func (this *BulletManager) Update() {
	for i := len(this.Bullets) - 1; i >= 0; i-- {
		bullet := this.Bullets[i]
		bullet.Update()
		if bullet.Dead {
			this.Bullets = append(this.Bullets[:i], this.Bullets[i+1:]...)
		}
	}
}
