package server

import (
	"github.com/gofiber/contrib/websocket"
)

func writeToClients(clients []*websocket.Conn, msg []byte) error{
	for _, client := range clients {
		if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
			client.Close()
			return err
		}
	}
	return nil
}