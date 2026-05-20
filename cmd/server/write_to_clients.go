package server

import (
	"broadcast-server/internal/model"
	"log"

	"github.com/gofiber/contrib/websocket"
)

func writeToClients(clients *model.Clients, msg []byte) error{
	clients.Mu.Lock()
	defer clients.Mu.Unlock()
	for _, client := range clients.WsConns {
		if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Printf("Failed to write message to client %s: %v", client.RemoteAddr(), err)
			return err
		}
	}
	return nil
}

func writeToClient(client *websocket.Conn, msg []byte) error{
	if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
		return err
	}
	return nil
}