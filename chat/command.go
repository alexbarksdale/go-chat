package chat

import (
	"fmt"
	"net"
)

// cmdId represents a command.
type cmdId int

const (
	CMD_JOIN cmdId = iota
	CMD_NAME
	CMD_MSG
	CMD_ROOMS
	CMD_LEAVE
)

// command represents a command.
type command struct {
	id   cmdId
	user *user
	args []string
}

// commandListener listens for commands on the server to execute.
func commandListener(s *server) {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_JOIN:
			cmd.join(s, cmd.args[0])
		case CMD_NAME:
		case CMD_ROOMS:
		case CMD_LEAVE:
			cmd.leave(s)
		}
	}
}

// join handles joining or creating a new room if one doesn't exist.
func (c *command) join(s *server, roomName string) {
	r, exist := s.rooms[roomName]
	if !exist {
		newRoom := fmt.Sprintf("Room doesn't exist, creating a new room: %s\n", c.args[0])
		c.user.sendMsg(newRoom)
		r = &room{
			name:  c.args[0],
			users: make(map[net.Addr]*user),
		}

		s.rooms[roomName] = r
	}

	c.leave(s)

	// Add user to the room
	r.users[c.user.conn.RemoteAddr()] = c.user

	// Assign the command executer's room
	c.user.room = r

	c.user.sendMsg(fmt.Sprintf("You joined the room: %s", r.name))
}

func (c *command) leave(s *server) {
	if c.user.room != nil {
		delete(s.rooms[c.user.room.name].users, c.user.conn.RemoteAddr())
	}
}
