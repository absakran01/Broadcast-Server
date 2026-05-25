package client

import (
	"log"
	"fmt"
	"strconv"
	"github.com/gorilla/websocket"
)

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