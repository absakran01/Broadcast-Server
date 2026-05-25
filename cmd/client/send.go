
package client

import (
	"bufio"
	"fmt"
	"os"
	"github.com/gorilla/websocket"
	"broadcast-server/internal/model"
)

func writeUserInput(clientConnection *model.ClientConnection, quit <-chan struct{}, reconnect chan<- error, connectionEstablished bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("connect to server and start typing messages (type 'exit' to quit):")

	for {
		select {
			case <-quit:
				return
			default:
				if connectionEstablished {
					go processInput(reader)
				}

				msgContent := <-inCh
				if msgContent == "" {
					continue
				}

				if msgContent == "ACK" {
					err := clientConnection.Conn.WriteMessage(websocket.TextMessage, []byte("ACK"))
					if err != nil {
						reconnect <- err
					}
					continue
				}
						
				publishMessage(clientConnection, msgContent)
			}
		
	}
}