package main

import (
	"broadcast-server/cmd/client"
	"broadcast-server/cmd/server"
	"flag"
)

func main() {
	port := flag.String("port", "8080", "Broadcast server port")
	host := flag.String("host", "127.0.0.1", "Host")
	mode := flag.String("mode", "client", "Mode (client or server)")
	flag.Parse()
	// redisAddr := flag.String("redis", "localhost:6379", "Redis server address")

	if *mode == "server" {
		server.Start(*host, *port)
		defer server.Shutdown()
	}

	if *mode == "client" {
		client.CreateSocketConnection(*host, *port)
	}

}
