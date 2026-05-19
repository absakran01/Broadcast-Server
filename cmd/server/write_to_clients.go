package server

import (
	"redisrelay/model"

	"github.com/gofiber/contrib/websocket"
)

func writeToClients(clients *model.Clients, msg []byte) error{
	clients.Mu.Lock()
	defer clients.Mu.Unlock()
	for _, client := range clients.WsConns {
		if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
			client.Close()
			return err
		}
	}
	return nil
}