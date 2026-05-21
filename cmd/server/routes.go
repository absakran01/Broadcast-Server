package server

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/ws", HandleInitWsConnection)
	app.Get("/ws", func(c *fiber.Ctx) error {
		return HandleWsConnection(c)
	})
}
