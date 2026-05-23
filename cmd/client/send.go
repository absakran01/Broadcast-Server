
package client

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/gorilla/websocket"
	"broadcast-server/internal/model"
)

// import UID generator
// import "your/package/path/client"

var (
	input string
	ackCh = make(chan bool)
)

func writeUserInput(clientConnection *model.ClientConnection, quit <-chan struct{}, reconnect chan<- error, connectionEstablished bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("connect to server and start typing messages (type 'exit' to quit):")

	for {
		select {
			case <-quit:
				log.Println("Write goroutine stopping...")
				return
			default:
				// input, _ := reader.ReadString('\n')
				// input = strings.TrimSpace(input)
				// if input == "exit" {
				// 	log.Println("Exit command received. Stopping client...")
				// 	return
				// }
				if connectionEstablished {
					go func() {
						input, _ = reader.ReadString('\n')
						input = strings.TrimSpace(input)
						if input == "exit" {
							log.Println("Exit command received. Stopping client...")
							os.Exit(0)
						}
						inCh <- input
					}()
				}

				msgContent := <-inCh
				if msgContent == "" {
					continue
				}
				switch msgContent {
					case "exit":
						log.Println("Exit command received. Stopping client...")
						os.Exit(0)
				    case "ACK":
						log.Println("ACK command is reserved. Please enter a different message.")
					
						err := clientConnection.Conn.WriteMessage(websocket.TextMessage, []byte("ACK"))
						if err != nil {
							log.Printf("Error sending ACK: %v", err)
						}
						continue
					default:
						// continue with normal message processing
						
				}
				// Generate UID for the message
				uid := GenerateUID()
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

					if ackReceived {
						log.Printf("Message with UID %s sent successfully", uid)
					}


					if !ackReceived {
						log.Printf("Failed to get ACK for UID %s", uid)
					}
				}
			}
		
	}
}
// Waits for an ACK message from the server and sends the result to ackCh
func waitForAck(clientConnection *model.ClientConnection) bool {
	for {
		_, msg, err := clientConnection.Conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading ACK: %v", err)
			return false
		}
		if string(msg) == "ACK" {
			log.Println("ACK received for sent message")
			return true
		}
	}
}