package server

import (
	"broadcast-server/internal/model"
	"fmt"
	"log"

	"github.com/ReneKroon/ttlcache"
	"github.com/gofiber/contrib/websocket"
)

func syncClientMessages(sync []byte, cache *ttlcache.Cache, c *websocket.Conn) error {
	ok, localMsgIndx := parseSyncMessage(sync)
	if !ok {
		return fmt.Errorf("invalid sync message format: %s", string(sync))
	} else {
		log.Printf("Client's local message index: %d", localMsgIndx)
		if localMsgIndx < cache.Count() && localMsgIndx > -1 {
			for i := localMsgIndx; i < cache.Count(); i++ {
				if msg, ok := cache.Get(fmt.Sprintf("%d", i)); ok {
					c.WriteMessage(websocket.TextMessage, msg.(*model.Message).Content)
				}
			}
		}
	}
	return nil
}
