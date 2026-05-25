package client

import (
	"log"
	"net/http"
	"time"
	"github.com/gorilla/websocket"
)

func dialWithRetry(addr string) *websocket.Conn {
	dialer := websocket.Dialer{
		HandshakeTimeout: http.DefaultClient.Timeout,
	}

	for {
		conn, _, err := dialer.Dial(addr, headers)
		if err != nil {
			log.Printf("failed to connect to %s: %v", addr, err)
			log.Println("Retrying ...")
			time.Sleep(2 * time.Second)
			continue
		}
		log.Printf("Successfully connected to %s", addr)
		return conn
	}
}