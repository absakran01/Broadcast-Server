package comms

import (
	"log"

	"github.com/gofiber/contrib/websocket"
)

const (
	ACK = "ACK"
)

func AckMsg(msg []byte, client *websocket.Conn) error {

	if msg != nil {
		err := client.WriteMessage(websocket.TextMessage, []byte(ACK))
		if err != nil {
			log.Println("ERROR:", err)
		}
	}
	return nil
}
