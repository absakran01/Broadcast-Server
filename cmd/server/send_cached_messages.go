package server

import (
	"broadcast-server/internal/model"
	"fmt"
	"github.com/ReneKroon/ttlcache"
	"github.com/gofiber/contrib/websocket"
)

func sendCachedMessages(localMsgIndx int, cache *ttlcache.Cache, c *websocket.Conn) error {
	for i := localMsgIndx; i < cache.Count(); i++ {
		if msg, ok := cache.Get(fmt.Sprintf("%d", i)); ok {
			err := writeWithAck(c, msg.(model.Message))
			if err != nil {
				return err
			}
		}
	}
	return nil
}