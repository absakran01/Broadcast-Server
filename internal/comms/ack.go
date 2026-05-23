package comms

import (
	"log"
	"sync"

	"broadcast-server/internal/model"

	"github.com/gofiber/contrib/websocket"
)

var (
	ACK = "ACK"
	mu = sync.Mutex{}
)

func AckMsg(msg *model.Message, client *websocket.Conn) error {
	mu.Lock()
	defer mu.Unlock()
	if msg.Content != nil {
		err := client.WriteMessage(websocket.TextMessage, []byte(ACK))
		if err != nil {
			log.Println("ERROR:", err)
		}
	}
	return nil
}
