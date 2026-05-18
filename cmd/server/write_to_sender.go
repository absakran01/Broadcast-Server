package server

import "github.com/gofiber/contrib/websocket"

func writeToSender(client *websocket.Conn, ack []byte) error {
	if err := client.WriteMessage(websocket.TextMessage, ack); err != nil {
		client.Close()
		return err
	}
	return nil
}