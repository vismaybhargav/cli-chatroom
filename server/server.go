package server

import "cli-chatroom/client"

type Server struct {
	clients []*client.Client
}
