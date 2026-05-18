package client

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func recieve(remoteConn *websocket.Conn) {
	for {
		_, msg, err := remoteConn.ReadMessage()
		if err != nil {
			panic("failed to read ack from server: " + err.Error())
		}
		fmt.Println("Received from server:", string(msg))

		err = remoteConn.WriteMessage(websocket.TextMessage, []byte("ack"))
		if err != nil {
			panic("failed to send ack to server: " + err.Error())
		}
	}
}