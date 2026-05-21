package server

import (
	"broadcast-server/internal/comms"
	"broadcast-server/internal/model"
	"broadcast-server/internal/util"
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
		msg := model.NewMessage(clientID, globalMsgIndx, msgContent)

		Indx, err := util.ExtractMsgIndxFromMsgId(msg.ID)
		if err != nil {
			log.Printf("Failed to extract message index from message ID: %v", err)
			continue
		}

		msgs[Indx] = msg
		if string(msg.Content) == "ACK" {
			//TODO: handle ACK from client
			continue
		}

		err = comms.AckMsg(msg, c)
		if err != nil {
			log.Println("ERROR:", err)
		}

		globalMsgIndx++

		err = writeToClients(clients, msg.Content, msg.ID)
		if err != nil {
			log.Println("ERROR:", err)
		}
	}
}
