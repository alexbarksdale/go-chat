package chat

import (
	"log"
	"net"
)

// server represents the chat server.
type server struct {
	rooms    map[string]*room
	commands chan command
}

// CreateServer creates a new chat server.
func CreateServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

// Run starts the server.
func (s *server) Run() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

	defer l.Close()

	// Listen for commands executed on the server
	go commandListener(s)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("Unable to accept connection:", err)
		}

		go createUser(conn, s)
	}
}
