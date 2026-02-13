package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/ltdlvr/websocket-chat/internal/client"
	"github.com/ltdlvr/websocket-chat/internal/hub"
	log "github.com/sirupsen/logrus"
)

func main() {
	//Создать хаб и запустить его в горутине
	h := hub.New()
	go h.Run()
	//Создать fiber сервер
	app := fiber.New()
	//Зарегистрировать роут через инстанс файбера
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		cl := &client.Client{
			Conn:    c,
			Receive: make(chan []byte),
		}
		h.Register <- cl
		go cl.WritePump()
		cl.ReadPump(h.Broadcast, h.Unregister)
	}))
	//Запустить сервер
	go func() {
		err := app.Listen(":3030")
		if err != nil {
			log.Fatal("Server failed to start: ", err)
		}
	}()
	//реализовать graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
	log.Info("Server stopped")
}
