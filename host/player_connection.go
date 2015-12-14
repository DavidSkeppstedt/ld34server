package host

import (
	"encoding/json"
	"github.com/DavidSkeppstedt/ld34server/game"
	"io"
	"log"
	"net"
)

type PlayerConnection struct {
	conn        net.Conn
	player      *game.Player
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

		data, err := this.ReciveInput()
		if err != nil {
			log.Println("Problem detected")
			break
		}
		this.ChangeState(data)
	}
}
func (this *PlayerConnection) ChangeState(state []byte) {
	switch state[0] {
	case 1:
		this.player.AngleInc(game.TurnSpeed)
	case 2:
		this.player.AngleInc(-game.TurnSpeed)
	case 3:
		this.player.Alive = true
	}

}

type PlayersPackage struct {
	Player  game.Player   `json:"player"`
	Enemies []game.Player `json:"enemies"`
	Bullets []game.Bullet `json:"bullets"`
}

func (this *PlayerConnection) SendPosition() {
	data := &PlayersPackage{*this.player, game.Pmanager.PlayersFiltered(this.player),
		game.Bmanager.Get()}
	this.jsonEncoder.Encode(data)
}
func (this *PlayerConnection) ReciveInput() ([]byte, error) {
	data := make([]byte, 1)
	_, err := this.conn.Read(data)
	if err == io.EOF {
		log.Println(err.Error())
		return []byte{}, err
	}
	return data, nil
}
