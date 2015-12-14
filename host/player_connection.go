package host

import (
	"encoding/json"
	"github.com/DavidSkeppstedt/ld34server/game"
	"github.com/DavidSkeppstedt/ld34server/game/player"
	"io"
	"log"
	"net"
)

type PlayerConnection struct {
	conn        net.Conn
	player      *player.Player
	jsonEncoder *(json.Encoder)
}

func (this *PlayerConnection) Play() {
	this.jsonEncoder = json.NewEncoder(this.conn)
	p := game.Pmanager.CreatePlayer()
	defer game.Pmanager.RemovePlayer(p)
	this.player = p

	log.Println("Send inital player position to server", p.Pos)

	for {
		this.SendPosition()

		msg := make([]byte, 1024)
		_, err := this.conn.Read(msg)
		if err == io.EOF {
			break
		}
	}
}

type PositionPackage struct {
	Player  player.Position   `json:"player_pos"`
	Enemies []player.Position `json:"other_pos"`
}

func (this *PlayerConnection) SendPosition() {
	data := &PositionPackage{this.player.Pos, game.Pmanager.PositionsFiltered(this.player)}
	this.jsonEncoder.Encode(data)
}
