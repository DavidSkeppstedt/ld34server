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
	this.jsonEncoder.Encode(p.Pos) //should write to the socket.. hopefully
	log.Println("Send inital player position to server", p.Pos)
	for {
		msg := make([]byte, 1024)
		_, err := this.conn.Read(msg)
		if err == io.EOF {
			break
		}
		log.Println(string(msg))
	}
}
