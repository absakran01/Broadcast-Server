package client

import (
	"broadcast-server/internal/model"
	"fmt"
	"log"
)

var (
	inCh = make(chan string)
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
				log.Printf("Receive limit of %d reached, stopping receive goroutine", globalMsgIndx)
				return
			}


			_, msg, err := clientConnection.Conn.ReadMessage()

			if err != nil {
				select {
				case reconnect <- err: // Try to send
					log.Println("Signaling reconnect from receive")
				default: // Already signaled by write goroutine
				}
				return
			}

			if string(msg) == "ACK" {
				log.Println("ACK received from server")
				ackCh <- true
				continue
			}
			// clientConnection.Mu.Lock()
			// err = clientConnection.Conn.WriteMessage(websocket.TextMessage, []byte("ACK"))
			// clientConnection.Mu.Unlock()

			fmt.Println(string(msg))
			*msgIdx++

		}
	}
}
