package main

import (
	"broadcast-server/cmd/client"
	"broadcast-server/cmd/server"
	"os"
	"flag"
)

func main() {
	port := flag.String("p", "8080", "Broadcast server port")
	host := flag.String("h", "127.0.0.1", "Host")
    if len(os.Args) < 2 {
        println("Usage: broadcast-server [start|connect]")
        return
    }
	flag.Parse()

	switch os.Args[1] {
	case "start":
		server.Start(*host, *port)
		defer server.Shutdown()
	

	case "connect":
		client.CreateSocketConnection(*host, *port)
	default:
		println("Unknown command. Usage: broadcast-server [start|connect]")
	}
}
