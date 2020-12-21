package chat

// cmdId represents a command.
type cmdId int

const (
	CMD_JOIN cmdId = iota
	CMD_NAME
	CMD_ROOMS
	CMD_QUIT
)

// command represents a command.
type command struct {
	id   cmdId
	user *user
	args []string
}
