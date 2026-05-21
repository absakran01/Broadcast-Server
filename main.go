package main

import (
	"broadcast-server/cmd/client"
	"broadcast-server/cmd/server"
	"flag"
	"github.com/gofiber/fiber/v2"
)


func main() {
	port := flag.String("port", "8080", "Broadcast server port")
	host := flag.String("host", "127.0.0.1", "Host")
	mode := flag.String("mode", "client", "Mode (client or server)")
	flag.Parse()
	// redisAddr := flag.String("redis", "localhost:6379", "Redis server address")

	if *mode == "server" {
		serverApp := fiber.New()

		server.Routes(serverApp)

		if err := serverApp.Listen(*host + ":" + *port); err != nil {
			panic(err)
		}
	}

	if *mode == "client" {
		client.CreateSocketConnection(*host, *port)
	}

}
