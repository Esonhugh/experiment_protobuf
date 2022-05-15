package _interface

import "proto_test/proto/message"

type Client struct {
	*MsgEngine
}

func InitClient() *Client {
	newClient := &Client{
		&MsgEngine{
			Reciver: make(chan message.ServerMessage, 1),
			Sender:  make(chan message.ServerMessage, 1),
		},
	}
	return newClient
}

func (c *Client) Connect(s *Server) {
	c.Sender = s.Reciver
	c.Reciver = s.Sender
}
