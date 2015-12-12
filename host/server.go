package host

import (
	_ "github.com/DavidSkeppstedt/ld34server/game/player"
	"log"
	"net"
)

var connections int

func ListenAndServe() {
	log.Println("Initializing the server")
	socket, err := net.Listen("tcp", ":7978")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Initialization complete. Listening...")
	for {
		con, err := socket.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}

		go handleConnection(con)
	}
}
func Inc(con net.Conn) {
	log.Println("Connection made from", con.RemoteAddr().String(), "connection #", connections)
	connections++
}

func Dec(con net.Conn) {
	connections--
	log.Println(con.RemoteAddr().String(), "disconnected, connection #", connections)
}
func handleConnection(con net.Conn) {
	defer con.Close()
	defer Dec(con)
	Inc(con)
	playerConn := &PlayerConnection{conn: con}
	playerConn.Play()
}
