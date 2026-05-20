package client

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func recieve(remoteConn *websocket.Conn, localMsgIndx int) {
	for {
		_, msg, err := remoteConn.ReadMessage()
		if err != nil {
			panic("failed to read message from server: " + err.Error())
		}

		if string(msg) == "ACK" {
			log.Println("Received ACK from server")
			log.Printf("local msg indx: %d", localMsgIndx)
			continue
		}

		fmt.Println("Received from server:", string(msg))
		fmt.Printf("localMsgIndx: %d\n", localMsgIndx)
		localMsgIndx++

		err = remoteConn.WriteMessage(websocket.TextMessage, []byte("ACK"))
		if err != nil {
			panic("failed to send ack to server: " + err.Error())
		}
	}
}
