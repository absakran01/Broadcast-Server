package model

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type Clients struct {
	WsConns map[string]*websocket.Conn
	Mu      sync.Mutex
}

func (clients *Clients) Add(c *websocket.Conn) {
	clients.Mu.Lock()
	defer clients.Mu.Unlock()
	clients.WsConns[c.RemoteAddr().String()] = c
}

func (clients *Clients) Remove(c *websocket.Conn) {
	clients.Mu.Lock()
	defer clients.Mu.Unlock()
	delete(clients.WsConns, c.RemoteAddr().String())
}
