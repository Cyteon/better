package router

import (
	"better/handlers"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
    api := app.Group("/_/")

    api.Get("/", func (c *fiber.Ctx) error { return c.SendString("OK") })

    messages := api.Group("/messages")
    messages.Get("/", handlers.HandleAllMessages)
    messages.Get("/:channel", handlers.HandleGetMessagesInChannel)
    messages.Post("/", handlers.HandleCreateMessage)
    messages.Put("/:id", handlers.HandleUpdateMessage)
}