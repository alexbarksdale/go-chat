package chat

import "net"

// room represents a chat room.
type room struct {
	name  string
	users map[net.Addr]*user
}

func (r *room) sendRoomMsg(msg string) {
	for _, user := range r.users {
		user.sendMsg(msg)
	}
}
