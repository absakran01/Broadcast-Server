package server

import (
	"broadcast-server/internal/comms"
	"broadcast-server/internal/model"
	"log"

	"github.com/gofiber/contrib/websocket"
)

func manageClientMessages(c *websocket.Conn, clients *model.Clients, clientID string) {
	for {
		_, msgContent, err := c.ReadMessage()
		if err != nil {
			log.Println("ERROR:", err)
			break
		}

		// Parse UID and content from the message
		msgUID, content, err := model.ParseUIDAndContent(msgContent)
		if err != nil {
			log.Printf("Invalid message format: %v", err)
			continue
		}
		msg := model.NewMessage(clientID, msgUID, cache.Count(), content)

		if string(msg.Content) == "ACK" {
			//TODO: handle ACK from client
			continue
		}
		cache.Set(msg.ID, msg)
		log.Printf("Received message from client %s: %s (ID: %s)", clientID, string(msg.Content), msg.ID)

		err = comms.AckMsg(msg, c)
		if err != nil {
			log.Println("ERROR:", err)
		}

		err = writeToClients(clients, msg.Content, msg.ID)
		if err != nil {
			log.Println("ERROR:", err)
		}
	}
}
