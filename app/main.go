package main

import (
	"log"
	"net"
)

func main() {
	log.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		log.Fatal("Failed to bind to port 6379")
	}
	conn, err := l.Accept()
	if err != nil {
		log.Fatal("Error accepting connection: ", err.Error())
	}
	var buf []byte
	_, err = conn.Read(buf)
	if err != nil {
		log.Fatal("Error reading the conn:", err)
	}
	response := []byte("+PONG\r\n")
	conn.Write(response)
	conn.Close()
}
