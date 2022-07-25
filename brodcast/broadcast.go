package brodcast

import "Chat/core/domain"

type Service struct {
	messages domain.Message
	incoming chan domain.Client
	leaving  chan domain.Client
}

func NewBroadcastService(messages domain.Message, incoming chan domain.Client, leaving chan domain.Client) *Service {
	return &Service{
		messages: messages,
		incoming: incoming,
		leaving:  leaving,
	}
}

func (s *Service) Broadcast() {
	clients := make(map[domain.Client]bool)
	for {
		select {
		case message := <-s.messages:
			for client := range clients {
				client <- message
			}
		case newClient := <-s.incoming:
			clients[newClient] = true
		case leavingClient := <-s.leaving:
			delete(clients, leavingClient)
			close(leavingClient)
		}
	}
}
