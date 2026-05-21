package client

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func receive(conn *websocket.Conn, quit <-chan struct{}, reconnect chan<- error, msgIdx *int) {
	for {
		select {
		case <-quit:
			log.Println("Receive goroutine stopping...")
			return
		default:

			_, msg, err := conn.ReadMessage()
			if err != nil {
				select {
				case reconnect <- err: // Try to send
					log.Println("Signaling reconnect from receive")
				default: // Already signaled by write goroutine
				}
				return
			}

			if string(msg) == "ACK" {
				continue
			}

			fmt.Println(string(msg))
			*msgIdx++

			conn.WriteMessage(websocket.TextMessage, []byte("ACK"))
		}
	}
}
