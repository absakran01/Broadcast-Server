package client

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

// import UID generator
// import "your/package/path/client"

var (
	input string
)

func writeUserInput(conn *websocket.Conn, quit <-chan struct{}, reconnect chan<- error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("connect to server and start typing messages (type 'exit' to quit):")

	for {
		select {
		case <-quit:
			log.Println("Write goroutine stopping...")
			return
		default:
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "exit" {
				log.Println("Exit command received. Stopping client...")
				return
			}

			// Generate UID for the message
			uid := GenerateUID()
			// Format: UID|message
			msgWithUID := fmt.Sprintf("%s|%s", uid, input)

			err := conn.WriteMessage(websocket.TextMessage, []byte(msgWithUID))
			if err != nil {
				select {
				case reconnect <- err: // Try to send
					log.Println("Signaling reconnect from write")
				default: // Already signaled by receive goroutine
				}
				return
			}
		}
	}
}
