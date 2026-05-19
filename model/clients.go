package model

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type Clients struct {
	WsConns map[string]*websocket.Conn
	Mu 	    sync.Mutex
}