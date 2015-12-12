package player

type Player struct {
	Pos   Position
	Angle float64
}

type Position struct {
	X float64
	Y float64
}

type PlayerManager struct {
	Players []Player
}

func (this *PlayerManager) CreatePlayer() Player {
	p := &Player{}
	this.Players = append(this.Players, *p)
	return *p
}
