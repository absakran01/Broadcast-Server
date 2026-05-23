package client

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
	"broadcast-server/internal/model"
)

func establishConnection(clientConnection *model.ClientConnection) (error) {
	for{
		globalMsgIndx, err := getMsgIndx(clientConnection.Conn)
		if err != nil {
			log.Printf("failed to get index: %v", err)
			clientConnection.Conn.Close()
			time.Sleep(2 * time.Second)
			continue
		}

		clientConnection.Conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("SYNC:%d", localMsgIndx)))
		if localMsgIndx  < 0 {
			localMsgIndx = globalMsgIndx
		} else{
			go writeUserInput(clientConnection, quit, reconnect, false)
			receive(clientConnection, quit, reconnect, &localMsgIndx, true)
		}

		_, connEst, err := clientConnection.Conn.ReadMessage()
		if err != nil || string(connEst) != "ACK_CONNECTION_ESTABLISHED" {
			return fmt.Errorf("failed to read connection ACK: %v", err)
		}

		err = clientConnection.Conn.WriteMessage(websocket.TextMessage, []byte("ACK_CONNECTION_ESTABLISHED"))
		if err != nil {
			return fmt.Errorf("failed to write connection ACK: %v", err)
		}


		return nil
	}
}