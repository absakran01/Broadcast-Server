package main

import (
	"flag"
	"broadcast-server/cmd/client"
	"broadcast-server/cmd/server"
	"broadcast-server/model"

	"github.com/gofiber/contrib/websocket"
	fiber "github.com/gofiber/fiber/v2"
)


var (
	clients = &model.Clients{
		WsConns: make(map[string]*websocket.Conn),
	}

)


func main() {
	port := flag.String("port", "8080", "Port")
    host := flag.String("host", "127.0.0.1", "Host")
	mode := flag.String("mode", "client", "Mode (client or server)")
	flag.Parse()
    // redisAddr := flag.String("redis", "localhost:6379", "Redis server address")

	if *mode == "server" {
		serverApp := fiber.New()

		server.Routes(serverApp, clients)

		if err := serverApp.Listen(*host + ":" + *port); err != nil {
			panic(err)
		}
	}

	if *mode == "client" {
		client.Connect(*host, *port)
	}

}

