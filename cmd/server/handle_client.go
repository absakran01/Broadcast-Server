package server

import (
	"broadcast-server/internal/model"
	"log"
	"strconv"
	"github.com/gofiber/contrib/websocket"
)

var (
	globalMsgIndx = 0
	msgs          = make(map[int]*model.Message)
		clients = &model.Clients{
		WsConns: make(map[string]*websocket.Conn),
	}
)

func HandleCLient() func(c *websocket.Conn) {
	return func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		log.Println("connection allowed:", c.Locals("allowed"))  // true

		clientID := model.GenClientID(c)
		uid := clients.Add(c, clientID)
		defer clients.Remove(uid)

		log.Printf("Client connected: %s (ID: %s)", c.RemoteAddr(), clientID)

		// Send the current global message index to the client for synchronization
		writeToClient(c, []byte(strconv.Itoa(globalMsgIndx)))
		_, sync, err := c.ReadMessage()
		if err != nil {
			log.Printf("Failed to read sync message from client: %v", err)
			return
		}

		// Sync client messages based on the received sync message and replay any missed messages
		err = syncClientMessages(sync, msgs, c)
		if err != nil {
			log.Printf("Failed to sync client messages: %v", err)
			return
		}

		manageClientMessages(c, clients, clientID)
	}
}
