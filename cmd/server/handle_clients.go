package server

import (
	"log"

	"github.com/gofiber/contrib/websocket"
)

const (
	ACK = "ack"
)

func HandleCLients(clients *[]*websocket.Conn) func(c *websocket.Conn) {
	return func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		log.Println(c.Locals("allowed"))  // true
		log.Println(c.Params("id"))       // 123
		log.Println(c.Query("v"))         // 1.0
		log.Println(c.Cookies("session")) // ""

		*clients = append(*clients, c)

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			msg []byte
			err error
		)
		for {
			if _, msg, err = c.ReadMessage(); err != nil {
				log.Println("ERROR:", err)
				break
			}
			if err := writeToSender(c, []byte(ACK)); err != nil {
				log.Println("ERROR:", err)
			}
			if err := writeToClients(*clients, msg); err != nil {
				log.Println("ERROR:", err)
			}

		}

	}
}