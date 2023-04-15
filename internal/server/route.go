package server

import (
	"github.com/Redish101/fiber-starter/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func initRoutes(app *fiber.App) {
	app.Get("/", handler.Home)
}