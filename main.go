package main

import (
	"github.com/alexbarksdale/go-chat/chat"
)

func main() {
	s := chat.CreateServer()
	s.Run()
}
