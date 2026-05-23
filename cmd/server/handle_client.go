package server

import (
	"broadcast-server/internal/model"
	"log"
	"github.com/gofiber/contrib/websocket"
)

func HandleCLient() func(c *websocket.Conn) {
	return func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		log.Println("connection allowed:", c.Locals("allowed"))  // true

		clientID := model.GenClientID(c)
		
		err := establishConnection(c, clientID)
		if err != nil {
			log.Printf("Failed to establish connection for client %s: %v", clientID, err)
			return
		}

		defer clients.Remove(clientID)

		log.Printf("Client connected: %s (ID: %s)", c.RemoteAddr(), clientID)

		manageClientMessages(c, clients, clientID)
	}
}