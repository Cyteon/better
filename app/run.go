package app

import (
	"os"

	"better/config"
	"better/database"
	"better/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run() error {
	err := config.LoadEnv()
	if err != nil {
		return err
	}

	err = database.StartDB()
	if err != nil {
		return err
	}

	defer database.CloseDB()

	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "${status} - ${method} ${path}\n",
	}))

    router.InitRoutes(app)

	config.AddSwaggerRoute(app)

	port := os.Getenv("PORT")
	app.Listen(":" + port)

	return nil
}
