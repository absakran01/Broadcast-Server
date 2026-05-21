package client

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

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

			err := conn.WriteMessage(websocket.TextMessage, []byte(input))
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
