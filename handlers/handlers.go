package handlers

import (
	"Chat/core/domain"
	"Chat/message"
	"bufio"
	"net"
)

type Handlers struct {
	messages   domain.Message
	incoming   chan domain.Client
	leaving    chan domain.Client
	msgService *message.Service
}

func NewHandlers(messages domain.Message, incoming chan domain.Client, leaving chan domain.Client, service *message.Service) *Handlers {
	return &Handlers{
		messages:   messages,
		incoming:   incoming,
		leaving:    leaving,
		msgService: service,
	}
}

func (h *Handlers) HandleConnection(conn net.Conn) {
	defer conn.Close()
	msg := make(chan string)
	go h.msgService.MessageWrite(conn, msg)

	clientName := conn.RemoteAddr().String()

	// Send just to the client his name
	msg <- "Welcome to the server, your name is " + clientName

	// Send the message to all the clients
	h.messages <- clientName + " has joined!"

	h.incoming <- msg

	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		h.messages <- clientName + ": " + inputMessage.Text()
	}

	// Remove the client from the list of clients
	h.leaving <- msg
	h.messages <- clientName + " has left!"
}
