package server

import (
	"log"
	"strconv"

	"github.com/gofiber/contrib/websocket"
)

func establishConnection(c *websocket.Conn) {
	// Send the current global message index to the client for synchronization
	writeToClient(c, []byte(strconv.Itoa(cache.Count())))
	_, sync, err := c.ReadMessage()
	if err != nil {
		log.Printf("Failed to read sync message from client: %v", err)
		return
	}

	// Sync client messages based on the received sync message and replay any missed messages
	err = syncClientMessages(sync, cache, c)
	if err != nil {
		log.Printf("Failed to sync client messages: %v", err)
		return
	}
}