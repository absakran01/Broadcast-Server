package comms

import (
	"log"

	"broadcast-server/internal/model"	
	"github.com/gofiber/contrib/websocket"
)

const (
	ACK = "ACK"
)

func AckMsg(msg *model.Message, client *websocket.Conn) error {

	if msg.Content != nil {
		err := client.WriteMessage(websocket.TextMessage, []byte(ACK))
		if err != nil {
			log.Println("ERROR:", err)
		}
	}
	return nil
}
