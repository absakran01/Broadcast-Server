package server

import (
	"broadcast-server/internal/model"
	"time"
	"github.com/ReneKroon/ttlcache"
	"github.com/gofiber/contrib/websocket"
)

var (
	clients       = &model.Clients{
		WsConns: make(map[string]*websocket.Conn),
	}
	// change as needed
	ttl = time.Minute * 3
	cache *ttlcache.Cache
)