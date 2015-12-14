package player

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

type Player struct {
	Pos   Position
	Angle float32
}

type Position struct {
	X float32 `json:"pos_x"`
	Y float32 `json:"pos_y"`
}

func (this *Player) Update() {
	this.move(0.09)
}
func (this *Player) AngleInc(amt float32) {
	this.Angle += amt
	if this.Angle > 360 {
		this.Angle = 0 + (this.Angle - 360)
	}
}

func (this *Player) move(speed float32) {
	this.Pos.X += speed * float32(math.Cos(float64(this.Angle)*math.Pi/180))
	this.Pos.Y += speed * float32(math.Sin(float64(this.Angle)*math.Pi/180))
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
	players := make([]*Player, len(this.Players)-1)
	for _, v := range this.Players {
		if v != toRemove {
			players = append(players, v)
		}
	}
	this.Players = players
}

func (this *PlayerManager) PositionsFiltered(toFilter *Player) []Position {

	allPos := make([]Position, len(this.Players)-1)
	if len(this.Players)-1 < 1 {
		return []Position{}
	}
	for _, v := range this.Players {
		if v != nil && v != toFilter {
			allPos = append(allPos, v.Pos)
		}

	}
	return allPos
}
