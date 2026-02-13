package hub

import (
	"github.com/ltdlvr/websocket-chat/internal/client"
)

type Hub struct {
	clients    map[*client.Client]bool
	Register   chan *client.Client
	Unregister chan *client.Client
	Broadcast  chan []byte
}

func New() *Hub {
	return &Hub{
		clients:    make(map[*client.Client]bool), //можно bool поменять на struct{} для экономии памяти
		Register:   make(chan *client.Client),
		Unregister: make(chan *client.Client),
		Broadcast:  make(chan []byte),
	}
}

func (h *Hub) Run() {
	for true {
		select {
		case c := <-h.Register:
			h.clients[c] = true
		case c := <-h.Unregister:
			delete(h.clients, c)
		case msg := <-h.Broadcast:
			for c := range h.clients {
				c.Send(msg)
			}
		}
	}
}
