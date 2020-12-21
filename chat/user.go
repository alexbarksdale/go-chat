package chat

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// user represents a user.
type user struct {
	conn     net.Conn
	name     string
	room     *room
	commands chan<- command
}

// createUser creates a new user.
func createUser(conn net.Conn, cmds chan command) {
	defer conn.Close()

	u := &user{
		conn:     conn,
		name:     "Default",
		commands: cmds,
	}

	u.readInput()
}

// readInput reads a user's input.
func (u *user) readInput() {
	for {
		data, err := bufio.NewReader(u.conn).ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read input", err)
			return
		}

		input := strings.Split(data, " ")
		cmd := input[0]

		switch cmd {
		case "/join":
			u.commands <- command{
				id:   CMD_JOIN,
				user: u,
				args: input[1:],
			}
		case "/name":
			u.commands <- command{
				id:   CMD_NAME,
				user: u,
				args: input[1:],
			}
		case "/rooms":
			u.commands <- command{
				id:   CMD_ROOMS,
				user: u,
				args: input[1:],
			}
		case "/leave":
			u.commands <- command{
				id:   CMD_LEAVE,
				user: u,
				args: input[1:],
			}
		}
	}
}
