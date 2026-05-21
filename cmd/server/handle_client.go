package server

import (
	"broadcast-server/internal/model"
	"log"
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

		//add client to clients to receive messages from other clients
		clientID := model.GenClientID(c)
		clients.Add(c, clientID)
		defer clients.Remove(clientID)

		establishConnection(c)

		log.Printf("Client connected: %s (ID: %s)", c.RemoteAddr(), clientID)

		manageClientMessages(c, clients, clientID)
	}
}



