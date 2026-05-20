package server

import (
	"broadcast-server/internal/comms"
	"broadcast-server/internal/model"
	"log"
	"strconv"

	"github.com/gofiber/contrib/websocket"
)


var (
	globalMsgIndx = 0
)

func HandleCLient(clients *model.Clients) func(c *websocket.Conn) {
	return func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		log.Println(c.Locals("allowed"))  // true
		log.Println(c.Params("id"))       // 123
		log.Println(c.Query("v"))         // 1.0
		log.Println(c.Cookies("session")) // ""

		clients.Add(c)
		defer clients.Remove(c)

		writeToClient(c, []byte(strconv.Itoa(globalMsgIndx)))

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			msg []byte
			err error
		)
		for {

			_, msg, err = c.ReadMessage()
			if err != nil {
				log.Println("ERROR:", err)
				break
			}

			err = comms.AckMsg(msg, c)
			if err != nil {
				log.Println("ERROR:", err)
			}

			globalMsgIndx++

			// err = writeToClients(clients, msg)
			// if err != nil {
			// 	log.Println("ERROR:", err)
			// }

		}

	}
}
