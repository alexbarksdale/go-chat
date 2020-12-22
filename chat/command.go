package chat

import (
	"fmt"
	"net"
	"strings"
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
			cmd.join(s)
		case CMD_NAME:
			cmd.setName()
		case CMD_MSG:
			cmd.sendMsg(s.rooms)
		case CMD_ROOMS:
			cmd.listRooms(s.rooms)
		case CMD_LEAVE:
			cmd.leaveRoom(s.rooms)
		}
	}
}

// join handles joining or creating a new room if one doesn't exist.
func (c *command) join(s *server) {
	roomName := c.args[0]

	r, exist := s.rooms[roomName]
	if !exist {
		newRoom := fmt.Sprintf("Room doesn't exist, creating a new room: %s\n", roomName)
		c.user.sendMsg(newRoom)
		r = &room{
			name:  c.args[0],
			users: make(map[net.Addr]*user),
		}

		s.rooms[roomName] = r
	}

	c.leaveRoom(s.rooms)

	// Add user to the room
	r.users[c.user.conn.RemoteAddr()] = c.user

	// Assign the user's current room
	c.user.room = r

	c.user.room.sendRoomMsg(fmt.Sprintf("%s joined the room.", c.user.name))
	c.user.sendMsg(fmt.Sprintf("You joined the room: %s\n", r.name))
}

// setName sets a custom display name.
func (c *command) setName() {
	c.user.name = c.args[0]
	c.user.sendMsg(fmt.Sprintf("Your name was set to: %s\n", c.user.name))
}

// sendMsg sends a message to the current room.
func (c *command) sendMsg(rooms map[string]*room) {
	if c.user.room == nil {
		c.user.sendMsg("You must join a room to send a message. View rooms with: /rooms")
		return
	}

	if len(c.args) <= 0 {
		c.user.sendMsg("Please provide a message.")
		return
	}

	msg := fmt.Sprintf("%s: %s", c.user.name, strings.Join(c.args[0:], " "))
	c.user.room.sendRoomMsg(msg)
}

// listRooms displays all available rooms to join.
func (c *command) listRooms(rooms map[string]*room) {
	output := make([]string, len(rooms))

	// Since we know the size, this is faster than appending.
	i := 0
	for room := range rooms {
		output[i] = room
		i++
	}

	c.user.sendMsg(fmt.Sprintf("Available rooms: %s", strings.Join(output, ", ")))
}

// leaveRoom leaves a room.
func (c *command) leaveRoom(rooms map[string]*room) {
	if c.user.room != nil {
		room := rooms[c.user.room.name]
		delete(room.users, c.user.conn.RemoteAddr())
		c.user.room.sendRoomMsg(fmt.Sprintf("👋 %s left the room.", c.user.name))
	}
}
