package host

import (
	"github.com/DavidSkeppstedt/ld34server/game/player"
	"io"
	"log"
	"net"
)

type PlayerConnection struct {
	conn   net.Conn
	player player.Player
}

func (this *PlayerConnection) Play() {
	for {
		msg := make([]byte, 1024)
		_, err := this.conn.Read(msg)
		if err == io.EOF {
			log.Println("Reading 0, dc")
			break
		}
		log.Println(string(msg))
		this.conn.Write([]byte("You said: " + string(msg) + "\n"))
	}
}
