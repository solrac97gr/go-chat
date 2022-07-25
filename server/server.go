package server

import (
	"Chat/brodcast"
	"Chat/core/domain"
	"Chat/handlers"
	"Chat/message"
	"fmt"
	"log"
	"net"
)

type Server struct {
	msgService       *message.Service
	broadcastService *brodcast.Service
	handlers         *handlers.Handlers
	incoming         chan domain.Client
	leaving          chan domain.Client
	messages         domain.Message
}

func NewServer() *Server {
	return &Server{
		incoming: make(chan domain.Client),
		leaving:  make(chan domain.Client),
		messages: make(domain.Message),
	}
}

func (s *Server) LoadServerComponents() {
	msgService := message.NewMessageService()
	hdl := handlers.NewHandlers(s.messages, s.incoming, s.leaving, msgService)
	broadcastService := brodcast.NewBroadcastService(s.messages, s.incoming, s.leaving)

	s.msgService = msgService
	s.handlers = hdl
	s.broadcastService = broadcastService
}

func (s *Server) Initialize(host *string, port *int) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}

	go s.broadcastService.Broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go s.handlers.HandleConnection(conn)
	}
}
