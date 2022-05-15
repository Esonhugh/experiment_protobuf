package _interface

import (
	"log"
	"proto_test/proto/message"
	"time"
)

type Message struct {
	*Server
	message.ServerMessage
}

type Handle func(*Message) error

type Server struct {
	*MsgEngine
	handlerFuncs []Handle
}

func InitSever() *Server {
	newServer := &Server{
		MsgEngine: &MsgEngine{
			Reciver: make(chan []byte, 1),
			Sender:  make(chan []byte, 1),
		},
	}
	return newServer
}

func (s *Server) Listen() {
	for true {
		for _, f := range s.handlerFuncs {
			each := &Message{
				s,
				s.Read(),
			}
			if err := f(each); err != nil {
				log.Println(err)
				break
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (s *Server) Register(funcA Handle) {
	s.handlerFuncs = append(s.handlerFuncs, funcA)
}
