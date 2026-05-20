package client

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gorilla/websocket"
)

var (
	input string
)

func write(remoteConn *websocket.Conn) {
	for {
		fmt.Print("Enter message to send: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ = reader.ReadString('\n')
		input = input[:len(input)-1] // Remove the newline character
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