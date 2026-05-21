package server

import (
	"broadcast-server/internal/model"
	"log"

	"github.com/gofiber/contrib/websocket"
)

func writeToClients(clients *model.Clients, msg []byte, msgID string) error{
	clients.Mu.Lock()
	payload := append([]byte(msgID+":"), msg...)
	defer clients.Mu.Unlock()
	for clientID, client := range clients.WsConns {
		if clientID == msgID[:len(clientID)] {
			continue
		}
		if err := client.WriteMessage(websocket.TextMessage, payload); err != nil {
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