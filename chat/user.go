package chat

import "net"

// user represents a user.
type user struct {
	conn     net.Conn
	name     string
	room     *room
	commands chan<- command
}

// createUser creates a new user.
func createUser(conn net.Conn, cmds chan command) *user {
	return &user{
		conn:     conn,
		name:     "Default",
		commands: cmds,
	}
}
