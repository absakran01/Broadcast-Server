package server

import (
	"fmt"
	"log"
	"strconv"
	"github.com/gofiber/contrib/websocket"
)

func establishConnection(c *websocket.Conn, clientID string) error {

	// Send the current global message index to the client for synchronization
	err := writeToClient(c, []byte(strconv.Itoa(cache.Count())))
	if err != nil {
		return err
	}

	_, sync, err := c.ReadMessage()
	if err != nil {
		return err
	}

	// Sync client messages based on the received sync message and replay any missed messages
	err = syncClientMessages(sync, cache, c)
	if err != nil {
		return err
	}

	err = writeToClient(c, []byte("ACK_CONNECTION_ESTABLISHED"))
	if err != nil {
		return err
	}

	_, ack, err := c.ReadMessage()
	if err != nil || string(ack) != "ACK_CONNECTION_ESTABLISHED" {
		return fmt.Errorf("something went wrong when receiving ack connection established from client")
	}
	log.Printf("Received ACK_CONNECTION_ESTABLISHED from client %s", clientID)

	//add client to clients to receive messages from other clients
	clients.Add(c, clientID)
	return nil
}