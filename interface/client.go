package _interface

type Client struct {
	*MsgEngine
}

func InitClient() *Client {
	newClient := &Client{
		&MsgEngine{
			Reciver: make(chan []byte, 1),
			Sender:  make(chan []byte, 1),
		},
	}
	return newClient
}

func (c *Client) Connect(s *Server) {
	c.Sender = s.Reciver
	c.Reciver = s.Sender
}
