package message

import (
	"fmt"
	"net"
)

type Service struct {
}

func NewMessageService() *Service {
	return &Service{}
}

func (s *Service) MessageWrite(conn net.Conn, messages <-chan string) {
	for msg := range messages {
		fmt.Fprintln(conn, msg)
	}
}
