package main

import (
	"flag"
	"log"
	"redisrelay/cmd/client"
	"redisrelay/cmd/server"

	"github.com/gofiber/contrib/websocket"
	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	port := flag.String("port", "8080", "Port")
    host := flag.String("host", "127.0.0.1", "Host")
	mode := flag.String("mode", "client", "Mode (client or server)")
	flag.Parse()
    // redisAddr := flag.String("redis", "localhost:6379", "Redis server address")


	if *mode == "server" {
		serverApp := fiber.New()
		clients := []*websocket.Conn{}

		serverApp.Get("/ws", handleInitWsConnection)
		serverApp.Get("/ws", func(c *fiber.Ctx) error {
			return handleWsConnection(c, &clients, *mode, *host, *port)
		})

		if err := serverApp.Listen(*host + ":" + *port); err != nil {
			panic(err)
		}
	}

	if *mode == "client" {
		client.Connect(*host, *port)
	}

}

func handleInitWsConnection(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}

	return fiber.ErrUpgradeRequired
}

func handleWsConnection(c *fiber.Ctx, clients *[]*websocket.Conn, mode string, host string, port string) error {

	return websocket.New((server.HandleCLients(clients)))(c)
}



func writeToClients(clients []*websocket.Conn, msg []byte) {
	for _, client := range clients {
		if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("write:", err)
			client.Close()
		}
	}
}
