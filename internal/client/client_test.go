package client

import (
	"testing"
)

func TestSendReceive(t *testing.T) {
	c := &Client{
		Receive: make(chan []byte, 1),
	}
	c.Send([]byte("Hello!!!"))
	if msg := <-c.Receive; string(msg) != "Hello!!!" {
		t.Errorf("Expected %q, got %q", "Hello!!!", string(msg))
	}

}
