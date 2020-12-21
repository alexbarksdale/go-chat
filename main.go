package main

import (
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("Unable to accept connection:", err)
		}

		go hello(conn)
	}
}

func hello(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("Hello"))
}
