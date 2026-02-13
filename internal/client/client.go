package client

import (
	"github.com/gofiber/websocket/v2"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	Conn    *websocket.Conn
	Receive chan []byte
}

func (c *Client) Send(msg []byte) {
	select {
	case c.Receive <- msg:

	default:

	}
}

func (c *Client) ReadPump(broadcastCh chan []byte, unregister chan *Client) {
	defer func() {
		unregister <- c
		close(c.Receive)
	}()
	for true {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Info("Client disconnected: ", err)
			return
		}
		broadcastCh <- msg
	}
}

func (c *Client) WritePump() {
	for msg := range c.Receive {
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Info("Client disconnected: ", err)
			return
		}
	}
}
