package client

import (
	"github.com/gorilla/websocket"
	"net/http"
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
	inCh = make(chan string)
	input string
	ackCh = make(chan bool)
)