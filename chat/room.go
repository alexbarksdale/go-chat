package chat

import "net"

// room represents a chat room.
type room struct {
	name  string
	users map[net.Addr]*user
}
