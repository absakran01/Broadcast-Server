package client

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var (
	headers = http.Header{
		"User-Agent": []string{"Broadcast-ServerClient/1.0"},
	}
	localMsgIndx = -1
)

func CreateSocketConnection(host string, port string) {
	addr := "ws://" + host + ":" + port + "/ws"

	for {

		quit := make(chan struct{})
		reconnect := make(chan error, 2)

		conn := dialWithRetry(addr)

		globalMsgIndx, err := getMsgIndx(conn)
		if err != nil {
			log.Printf("failed to get index: %v", err)
			conn.Close()
			time.Sleep(2 * time.Second)
			continue
		}

		if globalMsgIndx > localMsgIndx && localMsgIndx != -1 {
			conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%d", localMsgIndx)))
		} else {
			localMsgIndx = globalMsgIndx
		}

		go receive(conn, quit, reconnect, &localMsgIndx)
		go writeUserInput(conn, quit, reconnect)
		// go waitFor5SecsThenDisconnect(conn, quit, reconnect)

		err = <-reconnect
		log.Printf("Reconnection triggered: %v", err)

		close(quit)
		conn.Close()

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