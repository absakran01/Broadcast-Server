package client

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"broadcast-server/internal/model"
	"github.com/gorilla/websocket"
)

var (
	headers = http.Header{
		"User-Agent": []string{"Broadcast-ServerClient/1.0"},
	}
	localMsgIndx = -1
	globalMsgIndx = -1
	quit        chan struct{}
	reconnect   chan error
	conn        *websocket.Conn

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

func dialWithRetry(addr string) *websocket.Conn {
	dialer := websocket.Dialer{
		HandshakeTimeout: http.DefaultClient.Timeout,
	}

	for {
		conn, _, err := dialer.Dial(addr, headers)
		if err != nil {
			log.Printf("failed to connect to %s: %v", addr, err)
			log.Println("Retrying ...")
			time.Sleep(2 * time.Second)
			continue
		}
		log.Printf("Successfully connected to %s", addr)
		return conn
	}
}

func getMsgIndx(conn *websocket.Conn) (int, error) {
	_, msg, err := conn.ReadMessage()
	if err != nil {
		return 0, err
	}

	if len(msg) == 0 {
		log.Println("Received empty message for index, defaulting to 0")
		return 0, nil
	}

	localMsgIndx, err := strconv.Atoi(string(msg))
	if err != nil {
		return 0, fmt.Errorf("invalid message index '%s': %w", string(msg), err)
	}

	return localMsgIndx, nil
}

// func waitFor5SecsThenDisconnect(conn *websocket.Conn, quit chan struct{}, reconnect chan error) {
// 	time.Sleep(5 * time.Second)
// 	log.Println("Simulating disconnection after 5 seconds")
// 	reconnect <- fmt.Errorf("simulated disconnection")
// }