package host

import (
	"log"
	"net"
)

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
func handleConnection(con net.Conn) {
	defer con.Close()
	log.Println("A new connection was made!")
	con.Write([]byte("Hello"))
}
