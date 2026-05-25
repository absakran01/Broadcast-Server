package client

import (
	"log"
	"time"
	"broadcast-server/internal/model"
)

func CreateSocketConnection(host string, port string) {
	addr := "ws://" + host + ":" + port + "/ws"

	for {
		quit = make(chan struct{})
		reconnect = make(chan error, 2)

		conn = dialWithRetry(addr)

		clientConnection := &model.ClientConnection{Conn: conn}

		err := establishConnection(clientConnection)
		if err != nil {
			log.Printf("Failed to establish connection: %v", err)
			continue
		}

		go writeUserInput(clientConnection, quit, reconnect, true)
		go receive(clientConnection, quit, reconnect, &localMsgIndx, false)
	
		// go waitFor5SecsThenDisconnect(clientConnection, quit, reconnect)

		err = <-reconnect
		log.Printf("Reconnection triggered: %v", err)

		close(quit)
		clientConnection.Conn.Close()

		time.Sleep(2 * time.Second)
	}
}

// func waitFor5SecsThenDisconnect(conn *websocket.Conn, quit chan struct{}, reconnect chan error) {
// 	time.Sleep(5 * time.Second)
// 	log.Println("Simulating disconnection after 5 seconds")
// 	reconnect <- fmt.Errorf("simulated disconnection")
// }