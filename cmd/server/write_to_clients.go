package server

import (
	"broadcast-server/internal/model"

	"github.com/gofiber/contrib/websocket"
)

func writeToClients(clients *model.Clients, msg []byte) (string, error){
	clients.Mu.Lock()
	defer clients.Mu.Unlock()
	for _, client := range clients.WsConns {
		if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
			return client.IP(), err
		}
	}
	return "", nil
}

func writeToClient(client *websocket.Conn, msg []byte) error{
	if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
		return err
	}
	return nil
}