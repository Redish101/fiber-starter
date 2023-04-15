package server

import (
	"github.com/Redish101/fiber-starter/config"
	"github.com/gofiber/fiber/v2"
)


func Start() {
	app := fiber.New(fiber.Config{
		AppName: config.AppName,
		ServerHeader: config.AppName,
	})
	initRoutes(app)
	app.Listen(":3000")
}