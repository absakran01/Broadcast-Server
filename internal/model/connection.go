package model

import (
	gWebsocket "github.com/gorilla/websocket"
	"sync"
)

type ClientConnection struct {
	Conn *gWebsocket.Conn
	Mu  sync.Mutex
}