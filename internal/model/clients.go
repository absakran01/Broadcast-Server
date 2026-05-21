package model

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
	"broadcast-server/internal/util"
)

type Clients struct {
	WsConns map[string]*websocket.Conn
	Mu      sync.Mutex
}

func (clients *Clients) Add(c *websocket.Conn, clientID string) {
	clients.Mu.Lock()
	defer clients.Mu.Unlock()
	clients.WsConns[clientID] = c
}

func (clients *Clients) Remove(clientID string) {
	clients.Mu.Lock()
	defer clients.Mu.Unlock()
	delete(clients.WsConns, clientID)
}

func GenClientID(c *websocket.Conn) string {
	uid := util.GenerateUID()
	return c.RemoteAddr().String() + ":" + uid
}
