package server

import (
	"github.com/gofiber/contrib/websocket"
	fiber "github.com/gofiber/fiber/v2"
)

func HandleWsConnection(c *fiber.Ctx) error {
	return websocket.New((HandleCLient()))(c)
}
