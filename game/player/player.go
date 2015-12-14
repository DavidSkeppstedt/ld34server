package player

import (
	"log"
	"math"
	"math/rand"
	"sync"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))
var speed float32 = 1
var TurnSpeed float32 = 5
var width = 1280
var heigth = 720
var canShoot bool
var shootTimer = 500

//24 width 38 heigth
type Player struct {
	Pos   Position `json:"position"`
	Angle float32  `json:"angle"`
}

type Position struct {
	X float32 `json:"pos_x"`
	Y float32 `json:"pos_y"`
}

func (this *Player) Update() {
	this.move(speed)
	this.wrap()
	if shootTimer == 0 {
		shootTimer = 500
		if canShoot {
			log.Println("Shoot!")
		}
	} else {
		shootTimer -= 1
	}
	canShoot = true
}
func (this *Player) AngleInc(amt float32) {
	canShoot = false
	this.Angle += amt
	if this.Angle > 360 {
		this.Angle = 0 + (this.Angle - 360)
	}
}

func (this *Player) move(speed float32) {
	this.Pos.X += speed * float32(math.Cos(float64(this.Angle)*math.Pi/180))
	this.Pos.Y += speed * float32(math.Sin(float64(this.Angle)*math.Pi/180))
}

func (this *Player) wrap() {
	if this.Pos.X > float32(width) {
		this.Pos.X = 0
	}
	if this.Pos.X < 0 {
		this.Pos.X = float32(width)
	}

	if this.Pos.Y > float32(heigth) {
		this.Pos.Y = 0
	}
	if this.Pos.Y < 0 {
		this.Pos.Y = float32(heigth)
	}
}

type PlayerManager struct {
	Players []*Player
}

func (this *PlayerManager) Update() {
	for _, player := range this.Players {
		if player != nil {

			player.Update()
		}
	}
}

func (this *PlayerManager) CreatePlayer() *Player {
	pos := &Position{random.Float32() * 1280, random.Float32() * 720}
	p := &Player{Pos: *pos}
	this.Players = append(this.Players, p)
	return p
}

var listMutex sync.Mutex

func (this *PlayerManager) RemovePlayer(toRemove *Player) {
	listMutex.Lock()
	defer listMutex.Unlock()
	var players []*Player
	for _, v := range this.Players {
		if v != toRemove {
			players = append(players, v)
		}
	}
	this.Players = players
}

func (this *PlayerManager) PositionsFiltered(toFilter *Player) []Position {

	if len(this.Players)-1 < 1 {
		return []Position{}
	}
	var allPos []Position
	for _, v := range this.Players {
		if v != nil && v != toFilter {
			allPos = append(allPos, v.Pos)
		}

	}
	return allPos
}
func (this *PlayerManager) PlayersFiltered(toFilter *Player) []Player {

	if len(this.Players)-1 < 1 {
		return []Player{}
	}
	var allPlayers []Player
	for _, v := range this.Players {
		if v != nil && v != toFilter {
			allPlayers = append(allPlayers, *v)
		}

	}
	return allPlayers
}
