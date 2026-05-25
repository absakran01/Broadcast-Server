package server

import (
	"fmt"
	"github.com/ReneKroon/ttlcache"
	"github.com/gofiber/contrib/websocket"
)

func syncClientMessages(sync []byte, cache *ttlcache.Cache, c *websocket.Conn) error {
	ok, localMsgIndx := parseSyncMessage(sync)
	if !ok {
		return fmt.Errorf("invalid sync message format: %s", string(sync))
	} else {
		if localMsgIndx < cache.Count() && localMsgIndx > -1 {
			err := sendCachedMessages(localMsgIndx, cache, c)
			if err != nil {
				return err
			}
		}
	}
	return nil
}