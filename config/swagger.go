package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func AddSwaggerRoute(app *fiber.App) {
	app.Get("/docs/*", swagger.HandlerDefault)
}
