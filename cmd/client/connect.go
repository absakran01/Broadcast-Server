package client

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var (
	headers = http.Header{
		"User-Agent": []string{"RedisRelayClient/1.0"},
	}
	localMsgIndx = -1
)

func Connect(host string, port string) {
	addr := "ws://" + host + ":" + port + "/ws"

	dialer := websocket.Dialer{
		HandshakeTimeout: http.DefaultClient.Timeout,
	}
	remoteConn, _, err := dialer.Dial(addr, headers)
	if err != nil {
		log.Printf("failed to connect to %s: %v", addr, err)
		return
	}

	globalMsgIndx, err := getMsgIndx(remoteConn)
	if err != nil {
		log.Printf("failed to set index: %v", err)
		return
	}

	if localMsgIndx > globalMsgIndx && localMsgIndx != -1 {
		//TODO: request missed messages from server
	} else {
		localMsgIndx = globalMsgIndx
	}

	defer remoteConn.Close()

	

	go recieve(remoteConn, localMsgIndx)
	write(remoteConn)
}

func getMsgIndx(conn *websocket.Conn) (int, error) {
	_, msg, err := conn.ReadMessage()
	if err != nil {
		return 0, err
	}

	// Trim null bytes and whitespace
	msg = bytes.TrimSpace(bytes.Trim(msg, "\x00"))

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
