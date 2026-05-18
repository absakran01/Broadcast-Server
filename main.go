package main

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	server := fiber.New()
	clients := []*websocket.Conn{}

	server.Get("/ws", handleInitWsConnection)
	server.Get("/ws", func(c *fiber.Ctx) error {
		return handleWsConnection(c, &clients)
	})

	if err := server.Listen(":8080"); err != nil {
		panic(err)
	}

}

func handleInitWsConnection(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}

	return fiber.ErrUpgradeRequired
}

func handleWsConnection(c *fiber.Ctx, clients *[]*websocket.Conn) error {
	return websocket.New(func(c *websocket.Conn) {
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
				log.Println("read:", err)
				break
			}
			if string(msg) == "ping" {
				msg = []byte("pong")
				writeToClients(*clients, msg)
			} else {
				writeToClients(*clients, []byte("no"))
			}


		}

	})(c)
}


func writeToClients(clients []*websocket.Conn, msg []byte) {
	for _, client := range clients {
		if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("write:", err)
			client.Close()
		}
	}
}
