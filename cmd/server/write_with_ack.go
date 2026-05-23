package server

import (
	"broadcast-server/internal/model"
	"fmt"

	"github.com/gofiber/contrib/websocket"
)

func writeWithAck(c *websocket.Conn, message model.Message) error {
	if err := writeToClient(c, message.Content); err != nil {
		return err
	}

	_, ack, err := c.ReadMessage()
	if err != nil || string(ack) != "ACK" {
		return fmt.Errorf("something went wrong when receiving ack from client")
	}

	return nil
}