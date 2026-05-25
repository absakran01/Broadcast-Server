package client

import (
	"broadcast-server/internal/model"
	"fmt"
	"log"
)

func receive(clientConnection *model.ClientConnection, quit <-chan struct{}, reconnect chan<- error, msgIdx *int, limited bool) {
	for {
		select {
		case <-quit:
			log.Println("Receive goroutine stopping...")
			return
		default:

			// If limited is true, stop receiving after reaching globalMsgIndx for client synchronization
			if limited && *msgIdx == globalMsgIndx {
				return
			}

			_, msg, err := clientConnection.Conn.ReadMessage()

			if err != nil {
				select {
				case reconnect <- err: // Try to send
					log.Println("reconnecting...")
				default: // Already signaled by write goroutine
				}
				return
			}

			if string(msg) == "ACK" {
				ackCh <- true
				continue
			}

			fmt.Println(string(msg))
			*msgIdx++

		}
	}
}
