package server

import (
	"broadcast-server/internal/model"
	"log"
	"fmt"
	"github.com/gofiber/contrib/websocket"
)

func syncClientMessages(clientID string, sync []byte, msgs map[int]*model.Message, c *websocket.Conn) error {
	ok, localMsgIndx := parseSyncMessage(sync)
	if !ok {
		return fmt.Errorf("invalid sync message format: %s", string(sync))
	} else {
		log.Printf("Client's local message index: %d", localMsgIndx)
		if localMsgIndx < len(msgs) && localMsgIndx > -1 {
			for i := localMsgIndx; i < len(msgs); i++ {
				c.WriteMessage(websocket.TextMessage, msgs[i].Content)
			}
		}
	}
	return nil
}
