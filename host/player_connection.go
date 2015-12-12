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
	player      player.Player
	jsonEncoder *(json.Encoder)
}

func (this *PlayerConnection) Play() {
	this.jsonEncoder = json.NewEncoder(this.conn)
	p := game.Pmanager.CreatePlayer()
	this.player = p

	log.Println("Send inital player position to server", p.Pos)

	for {
		this.SendPosition()

		msg := make([]byte, 1024)
		_, err := this.conn.Read(msg)
		if err == io.EOF {
			break
		}
		log.Println(string(msg))
	}
}

func (this *PlayerConnection) SendPosition() {
	this.jsonEncoder.Encode(this.player.Pos)
}
