package client

import (
	"fmt"
	"github.com/gorilla/websocket"
)

func write(remoteConn *websocket.Conn) {
	for {
		fmt.Scanln(&input)
		remoteConn.WriteMessage(websocket.TextMessage, []byte(input))
		_, ack, err := remoteConn.ReadMessage()
		if err != nil {
			panic("failed to read ack from server: " + err.Error())
		}
		if string(ack) != "ack" {
			panic("unexpected response from server: " + string(ack))
		}
	}
}