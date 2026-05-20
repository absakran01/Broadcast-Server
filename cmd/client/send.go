package client

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var (
	input string
)

func write(remoteConn *websocket.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message to send: ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		remoteConn.WriteMessage(websocket.TextMessage, []byte(input))
		time.Sleep(1 * time.Second) // Add a small delay to display ack before next input
	}
}
