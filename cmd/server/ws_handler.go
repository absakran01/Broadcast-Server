package server

import (
	"broadcast-server/internal/model"

	"github.com/gofiber/contrib/websocket"
	fiber "github.com/gofiber/fiber/v2"
)

func HandleWsConnection(c *fiber.Ctx, clients *model.Clients) error {
	return websocket.New((HandleCLient(clients)))(c)
}
