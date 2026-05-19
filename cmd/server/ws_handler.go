package server

import (
	"redisrelay/model"

	"github.com/gofiber/contrib/websocket"
	fiber "github.com/gofiber/fiber/v2"
)

func HandleWsConnection(c *fiber.Ctx, clients *model.Clients) error {
	return websocket.New((HandleCLients(clients)))(c)
}
