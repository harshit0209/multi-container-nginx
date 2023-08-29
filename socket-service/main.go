package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for WebSocket
	},
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/ws", hello)

	port := ":5000"
	log.Fatal(e.Start(port))
}

func hello(c echo.Context) error {
	ws, err := Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Read message from client
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Send the received message back to the client
		err = ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Error sending message:", err)
			break
		}
	}

	return nil
}
