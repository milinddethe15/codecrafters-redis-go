package main

import (
	"log"
	"net"
)

func main() {
	log.Println("Logs from redis server will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		log.Fatal("Failed to bind to port 6379")
	}
	conn, err := l.Accept()
	if err != nil {
		log.Fatal("Error accepting connection: ", err.Error())
	}
	go handleConn(conn)
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			log.Println("Error reading the conn: ", err)
			break
		}
		response := []byte("+PONG\r\n")
		conn.Write(response)
	}
}
