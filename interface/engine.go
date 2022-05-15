package _interface

import (
	"log"
	"proto_test/proto/message"
)

import "github.com/golang/protobuf/proto"

type MsgEngine struct {
	Reciver chan []byte
	Sender  chan []byte
}

func (m *MsgEngine) Send(data message.ServerMessage) *MsgEngine {
	var send, err = proto.Marshal(&data)
	if err != nil {
		log.Println("marshal error:", err)
	}
	log.Println("need send:", send)
	m.Sender <- send
	return m
}

func (m *MsgEngine) Read() message.ServerMessage {
	var msgs message.ServerMessage
	received := <-m.Reciver
	log.Println("received:", received)
	err := proto.Unmarshal(received, &msgs)
	if err != nil {
		log.Println("unmarshal error:", err)
	}
	return msgs
}
