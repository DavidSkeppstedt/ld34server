package host

import (
	"github.com/DavidSkeppstedt/ld34server/game/player"
	"net"
)

type PlayerConnection struct {
	conn   net.Conn
	player player.Player
}
