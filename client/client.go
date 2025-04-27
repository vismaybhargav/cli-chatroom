package client

import (
	"cli-chatroom/server"
	"golang.org/x/net/websocket"
)

type Client struct {
	name       string
	server     *server.Server
	connection *websocket.Conn
}

func NewClient(name string, connection *websocket.Conn, server *server.Server) *Client {
	return &Client{
		name:       name,
		connection: connection,
		server:     server,
	}
}
