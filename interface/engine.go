package _interface

import "proto_test/proto/message"

type MsgEngine struct {
	Reciver chan message.ServerMessage
	Sender  chan message.ServerMessage
}

func (m *MsgEngine) Send(data message.ServerMessage) error {
	m.Sender <- data
	return nil
}

func (m *MsgEngine) Read() message.ServerMessage {
	msg := <-m.Reciver
	return msg
}
