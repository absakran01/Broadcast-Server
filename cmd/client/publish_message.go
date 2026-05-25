package client

import (
	"broadcast-server/internal/model"
	"broadcast-server/internal/util"
	"fmt"
	"log"
	"github.com/gorilla/websocket"
)

func publishMessage(clientConnection *model.ClientConnection, msgContent string) {
	// Generate UID for the message
	uid := util.GenerateUID()
	msgWithUID := fmt.Sprintf("%s|%s", uid, msgContent)

	ackReceived := false
	for !ackReceived {
		clientConnection.Mu.Lock()
		err := clientConnection.Conn.WriteMessage(websocket.TextMessage, []byte(msgWithUID))
		clientConnection.Mu.Unlock()
		if err != nil {
			select {
			case reconnect <- err:
				log.Println("Signaling reconnect from write")
				default:
			}
			return
		}
				
		ackReceived = <-ackCh

		if !ackReceived {
			log.Printf("Failed to get ACK for UID %s", uid)
		}
	}
}