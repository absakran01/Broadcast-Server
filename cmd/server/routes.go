package server

import (

	"broadcast-server/model"

	fiber "github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, clients *model.Clients) {
	app.Get("/ws", HandleInitWsConnection)
	app.Get("/ws", func(c *fiber.Ctx) error {
		return HandleWsConnection(c, clients)
	})
}
