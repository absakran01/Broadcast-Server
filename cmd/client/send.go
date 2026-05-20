package client

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var (
	input string
)

func write(conn *websocket.Conn, quit <-chan struct{}, reconnect chan<- error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		select {
		case <-quit:
			log.Println("Write goroutine stopping...")
			return
		default:
			fmt.Print("Enter message: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			err := conn.WriteMessage(websocket.TextMessage, []byte(input))
			if err != nil {
				select {
				case reconnect <- err: // Try to send
					log.Println("Signaling reconnect from write")
				default: // Already signaled by receive goroutine
				}
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}
