package main

import (
	"log"
	_interface "proto_test/interface"
	"proto_test/proto/message"
	"time"
)

func handler(data *_interface.Message) error {
	log.Printf("Code: %v, Type: %v,Summary: %s,Content: %s", data.Code, data.Msgtype, data.Summary, data.Content)
	data.Server.Send(message.ServerMessage{
		Code:    20000 + data.Code,
		Msgtype: message.ServerMessage_SHORT,
		Summary: "200 OK",
		Content: "Server Recived! Your Content is " + data.Content,
	})
	return nil
}

func main() {
	s := _interface.InitSever()
	s.Register(handler)
	c := _interface.InitClient()
	c.Connect(s)
	go s.Listen()
	var code int64 = 1
	for true {
		data := c.Send(message.ServerMessage{
			Code:    code,
			Msgtype: message.ServerMessage_SHORT,
			Summary: "test",
			Content: "TEST HelloWorld!",
		}).Read()
		log.Printf("Code: %v, Type: %v,Summary: %s,Content: %s", data.Code, data.Msgtype, data.Summary, data.Content)
		code++
		time.Sleep(1 * time.Millisecond)
	}

}
