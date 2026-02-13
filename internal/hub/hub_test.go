package hub

import (
	"testing"
	"time"

	"github.com/ltdlvr/websocket-chat/internal/client"
)

func TestNew(t *testing.T) {
	h := New()

	if h == nil {
		t.Fatal("Hub is nil") //надо поучиться ошибки описывать
	}

	tests := []struct {
		name  string
		value any //это алиас для interface{}
	}{
		{"clients", h.clients},
		{"Register", h.Register},
		{"Unregister", h.Unregister},
		{"Broadcast", h.Broadcast},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.value == nil {
				t.Errorf("%v is nil", test.name)
			}
		})
	}
}
func TestRegister(t *testing.T) {
	h := New()
	go h.Run()
	c := &client.Client{}
	h.Register <- c
	time.Sleep(10 * time.Millisecond) //TODO Чем можно заменить?
	if h.clients[c] != true {
		t.Error("Client not registred")
	}
}
