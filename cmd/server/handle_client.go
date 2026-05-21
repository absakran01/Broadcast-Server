package server

import (
	"broadcast-server/internal/comms"
	"broadcast-server/internal/model"
	"broadcast-server/internal/util"
	"log"
	"strconv"

	"github.com/gofiber/contrib/websocket"
)

var (
	globalMsgIndx = 0
	msgs          = make(map[int]*model.Message)
)

func HandleCLient(clients *model.Clients) func(c *websocket.Conn) {
	return func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		log.Println(c.Locals("allowed"))  // true
		log.Println(c.Params("id"))       // 123
		log.Println(c.Query("v"))         // 1.0
		log.Println(c.Cookies("session")) // ""

		clientID := model.GenClientID(c)
		uid := clients.Add(c, clientID)
		defer clients.Remove(uid)

		writeToClient(c, []byte(strconv.Itoa(globalMsgIndx)))
		_, sync, err := c.ReadMessage()
		ok, localMsgIndx := parseSyncMessage(sync)
		if !ok || err != nil {
			log.Printf("Failed to read client's local message index: %v", err)
		} else {
			log.Printf("Client's local message index: %d", localMsgIndx)
			if localMsgIndx < len(msgs) && localMsgIndx > -1 {
				for i := localMsgIndx; i < len(msgs); i++ {
					c.WriteMessage(websocket.TextMessage, msgs[i].Content)
				}
			}

		}

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
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
}
