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

func (clients *Clients) Add(c *websocket.Conn, clientID string) string {
	clients.Mu.Lock()
	defer clients.Mu.Unlock()
	clients.WsConns[clientID] = c
	return clientID
}

func (clients *Clients) Remove(uid string) {
	clients.Mu.Lock()
	defer clients.Mu.Unlock()
	delete(clients.WsConns, uid)
}

func GenClientID(c *websocket.Conn) string {
	uid := util.GenerateUID()
	return c.RemoteAddr().String() + ":" + uid
}
